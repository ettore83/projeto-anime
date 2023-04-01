<script >
import { ref } from "vue";

const newAnime = ref("")


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

          
          <label for="tentei" class="label-inp">Anime name</label>
          <input id="tentei"
          class="imp-create"
          type="text"
          placeholder="Name of Anime">

          
          <label for="tentei" class="label-inp">Release date</label>
          <input
          class="imp-create"
          type="number"
          placeholder="When it was released">

        
          <label for="tentei" class="label-inp">Quantity episodios</label>
          <input
          class="imp-create"
          type="text"
          placeholder="Episodios released">

          
          <label for="tentei" class="label-inp">Genre type</label>
          <input
          class="imp-create"
          type="text"
          placeholder="Genre">

          
          <label for="tentei" class="label-inp">Written by</label>
          <input
          class="imp-create"
          type="text"
          placeholder="written by">

          
          <label for="tentei" class="label-inp">Description</label>
          <input
          class="imp-create"
          type="text"
          placeholder="Description">
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

#form-create{
  display: grid;
  gap: 2px;
}

.label-inp{
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

@media (min-width:700px){
 .about{
  text-align: left;
 }

  .imp-create{
  width: 30%;
  text-align: left;
  
}

}
        

  
 
 
  </style>