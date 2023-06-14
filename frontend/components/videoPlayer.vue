<script>
    let playerWidth = 640;
    let playerHeight = 360;
    let source;

	onMounted(()=>{
        source = new MediaSource();

        source.addEventListener('sourceopen',(ev)=>{
            console.log("source open")
            let b = ev.target;
            
            //will need a way to dynamically set the codec but that is a later problem
            sourceBuffer = b.addSourceBuffer("audio/webm;codecs=opus")
            sourceBuffer.addEventListener('update',()=>{
                console.log("source buffer update")
                if (queue.length > 0 && !sourceBuffer.updating) {
                    sourceBuffer.appendBuffer(queue.shift());
                }
            })
            speakerReady = true;
        },false)
    });

</script>
<template>
    still here
    <video :width="{playerWidth}" :height="{playerHeight}">
        <source src="https://youtu.be/fx2Z5ZD_Rbo" type="video/youtube">
        Your browser does not support video players
    </video>
 </template>