<script setup >

</script>

<template>
    <main>
      <div class="about">
        <h1>Create a new Anime</h1>
      
        <form action="submit" id="form-create">
          <p>Anime name</p>
          <input id="tentei"
          class="imp-create"
          type="text"
          placeholder="Name of Anime">
          <p>Release date</p>
          <input
          class="imp-create"
          type="number"
          placeholder="When it was released">
          <p>Quantity episodios</p>
          <input
          class="imp-create"
          type="text"
          placeholder="Episodios released">
          <p>Genre type</p>
          <input
          class="imp-create"
          type="text"
          placeholder="Genre">
          <p>Written by</p>
          <input
          class="imp-create"
          type="text"
          placeholder="written by">
          <p>Description</p>
          <input
          class="imp-create"
          type="text"
          placeholder="Description"
          >
          <br>
          <button id="btn-post">
          POST
          </button>
        </form>
      </div>
    </main>
  </template>
  
  <style >
.about{
  width: 100%;
  text-align: center;
  background-color:red;
  height: 900px;
}
h1{
  margin: auto;
}
p{
  color: aliceblue;
  width: 80%;
  margin: auto;
  text-align: left;
  margin-top: 20px;
  
}
.imp-create{
    width: 80%;
    padding: 15px;
    border-radius: 10px;
    margin: auto;
}  
        
 button{
  width: 80%
  padding :15px;
  border-radius:10px;
  margin:auto;
}       

  
 
 
  </style>