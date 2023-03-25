<script >
import { ref } from "vue";


</script>

<template>
    <main>
      <div class="about">
        <div id="bt-back">
          <RouterLink to="/">
          <img id="bt-img-back" src="../assets/ic_back_white@3x.png" alt=""></RouterLink>
          <h1>Create a new Anime</h1>
        </div>
      
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
          v-bind:disabled="realeseData < 4"
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

#bt-back{
  display: flex;
  width: 80%;
  margin: auto;
  align-items: end;
}

#bt-img-back{
  width: 30px;
  height: 30px;
  
}


.about{
  width: 100%;
  text-align: center;
  background-color:#344960;
  height: 900px;
}
h1{
  margin: auto;
  margin-top: 20px;
  color: aliceblue;
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
    border: none;
}  

#btn-post{
  width: 80%;
  padding :15px;
  border-radius:10px;
  margin:auto;
  margin-top: 40px;
  cursor: pointer;

}       

        

  
 
 
  </style>