<!-- NEVER HAVE A useState OUTSIDE OF A <script setup> It will cause a memory leak and will be shared across all instances of the website -->
<script setup>
    //This will start counting at zero and every time the button is clicked it will add one
    const counter = useState('counter', ()=> 0)
</script>
<template>
    <button @click="counter++">
        state counter
    </button>
    {{ counter }}
    <div>
      test <br/>

      <example/>
      
      <NuxtLink to="/page1">routing example</NuxtLink>
    </div>
</template>