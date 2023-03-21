<script setup >

</script>

<template>


    <main>
      <div class="about">
        <h1>Create a new Anime</h1>
      
        <form action="submit" id="form-create">
          <input id="tentei"
          class="imp-create"
          type="text"
          placeholder="Name of Anime">
          <input
          class="imp-create"
          type="number"
          placeholder="When it was released">
          <input
          class="imp-create"
          type="text"
          placeholder="Episodios released">
          <input
          class="imp-create"
          type="text"
          placeholder="Genre">
          <input
          class="imp-create"
          type="text"
          placeholder="written by">
          <input
          class="imp-create"
          type="text"
          placeholder="Description"
          >
        </form>
      </div>
    </main>
  </template>
  
  <style >
.about{
  width: 100%;
  text-align: center;
  background-color:var(--color-background);
}
h1{
  margin: auto;
}
.imp-create{
    width: 80%;
    padding: 10px;
    border-radius: 10px;
    margin: auto;
    
  }
 
 
  </style>