<!-- 
    Any file or directory in the pages  directory will become the route so the file path pages/page1.vue becomes the route /page1
    and the file path pages/examples/dumb.vue becomes the route /example/dumb
    
 -->

<template>
    <header>
        page 1
    </header>
    <NuxtLink to="/page2">To page 2</NuxtLink>
        
  </template>