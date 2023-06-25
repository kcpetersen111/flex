<script setup>
    let sourceBuffer;
    //this is temporary just to pass the check, We will want this to be set dynamically
    let codec = 'video/mp4; codecs="avc1.4D401F"';
    //eventually this will be used to put data in while the source buffer is not ready
    let queue = [];
 
	onMounted(()=>{
        if(!MediaSource.isTypeSupported(codec)){
            console.log(`codec: ${codec} is not supported`);
            return;
        }
        sourceBuffer = new MediaSource();

        sourceBuffer.addEventListener('sourceopen',(ev)=>{
            console.log("source open");
            let b = ev.target;
            
            sourceBuffer = b.addSourceBuffer(codec);
            sourceBuffer.addEventListener('update',()=>{
                if (queue.length > 0 && !sourceBuffer.updating) {
                    sourceBuffer.appendBuffer(queue.shift());
                }
            });
        },false)
    });

</script>
<script>
    import {serverAddr} from "~/utils/variables";
    
    function until(conditionFunction) {
        const poll = resolve => {
            if(conditionFunction()) resolve();
            else setTimeout(_ => poll(resolve), 400);
        }
        return new Promise(poll);
    }

    async function startWebSocket() {
        var doneConnecting = false;
        console.log("starting websocket")
        const ws = new WebSocket(`ws://${serverAddr}/ws`);
        ws.addEventListener("connect",_=>{
            console.log("web socket connection is started");
        });
        ws.addEventListener("open", _=>{
            console.log("ws open")
            doneConnecting = true;
        });
        ws.addEventListener("message",messageEvent=>{
            console.log(JSON.parse(messageEvent.data))
        })

        await until(_=>doneConnecting==true);
        return ws;
    }
    async function pingServer() {
        const ws = await startWebSocket();
        ws.send(JSON.stringify({"Test":"Success"}));
    }
</script>
<template>
    <button v-on:click="pingServer">start ws</button>
    
 </template>