<!-- NEVER HAVE A useState OUTSIDE OF A <script setup> It will cause a memory leak and will be shared across all instances of the website -->
<script setup>
    //This will start counting at zero and every time the button is clicked it will add one
    const counter = useState('counter', ()=> 0)
    const word = useState('word', ()=>"ahh")
</script>
<template>
    <!-- state example 1 -->
    <button @click="counter++">
        state counter
    </button>
    {{ counter }}
    <br/>

<!-- state example two -->
      <button @click="()=>{word='a'+word}">
        state counter
    </button>
    {{ word }}
    <br/>

    <div>
        <!-- This is a component  -->
      <example/>

      <!-- This is a routing example -->
      <NuxtLink to="/page1">routing example</NuxtLink>
    </div>
</template>