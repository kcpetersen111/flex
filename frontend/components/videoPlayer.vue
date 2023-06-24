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
    

    function startWebSocket() {
        console.log("starting websocket")
        const ws = new WebSocket(`ws://${serverAddr}/ws`);
        ws.addEventListener("connect",()=>{
            console.log("web socket connection is started");
        })
        return ws;
    }
</script>
<template>
    <button v-on:click="startWebSocket">start ws</button>
    
 </template>