<script >
import { ref } from "vue";

const animeName = ref("")
const releaseDate = ref("")
const quantityEpisodios = ref("")
const genreType = ref("")
const writtenBy = ref("")
const description = ref("")

const newAnime = ref("")

const nameIsValid = !!animeName.value
const releasedIsValid = !!releaseDate.value
const episodiosIsValid = !!quantityEpisodios.value
const genreIsValid = !!genreType.value
const writtenIsValid = !!writtenBy.value
const descriptionIsValid = !!description.value 





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

          
          <label for="anime-name" class="label-inp">Anime name</label>
          <input id="tentei"
          v-model = "animeName"
          class="imp-create"
          type="text"
          placeholder="Name of Anime">
          <p v-if="!nameIsValid" class="error-msg">The anime name fild is required</p>

          
          <label for="release-date" class="label-inp">Release date</label>
          <input
          v-model = "releaseDate"
          class="imp-create"
          type="number"
          placeholder="When it was released">
          <p v-if="!releasedIsValid" class="error-msg">The release date fild is required</p>

        
          <label for="quantity-episodios" class="label-inp">Quantity episodios</label>
          <input
          v-model = "quantityEpisodios"
          class="imp-create"
          type="text"
          placeholder="Episodios released">
          <p v-if="!episodiosIsValid" class="error-msg">The quantity episodios fild is required</p>

          
          <label for="genre-type" class="label-inp">Genre type</label>
          <input
          v-model = "genreType"
          class="imp-create"
          type="text"
          placeholder="Genre">
          <p v-if="!genreIsValid" class="error-msg">The genre type fild is required</p>

          
          <label for="written-by" class="label-inp">Written by</label>
          <input
          v-model = "writtenBy"
          class="imp-create"
          type="text"
          placeholder="written by">
          <p v-if="!writtenIsValid" class="error-msg">The written by fild is required</p>

          
          <label for="description" class="label-inp">Description</label>
          <input
          v-model = "description"
          class="imp-create"
          type="text"
          placeholder="Description">
          <p v-if="!descriptionIsValid" class="error-msg">The description fild is required</p>
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