<script setup >

import { reactive } from "vue";
import { useRouter } from "vue-router";
import { required, numeric, integer, minLength } from "@vuelidate/validators";
import { useVuelidate } from "@vuelidate/core";
import criar from "../services/config";

    const MyAnime = reactive({
      animeName : '',
      releaseDate : '',
      quantityEpisodios : '',
      genreType : '',
      writtenBy : '',
      description : '',
    });
    console.log(MyAnime.animeName)

    const rules = {
      animeName: { required },
      releaseDate: { required, integer, minLength: minLength(8) },
      quantityEpisodios: { required, numeric},
      genreType: { required },
      writtenBy: { required },
      description: { required },
    };         

    const v$ = useVuelidate (rules, MyAnime);    

    const sendForm = async () => {
      MyAnime.loading = true;
      const isFormCorrect = await v$.value.$validate();

      if (!isFormCorrect) {
        MyAnime.loading = false;
        return;
      }

      criar
        .insertData({
          name: MyAnime.animeName,
          release_date: MyAnime.releaseDate,
          episodes: MyAnime.quantityEpisodios,
          genre: MyAnime.genreType,
          author: MyAnime.writtenBy,
          description: MyAnime.description,
        
        })
      .then(async (response) => {
        if (response.status != 200) {
          alert("Error: " + response.data);
          return;
        }

        console.log(response)
      })
      .catch((error) => {
        console.log(error);
      });
    MyAnime.loading = false;
};



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
          
      
          <input
            id="tentei"
            class="imp-create"
            type="text"
            v-model.trim="MyAnime.animeName"
            placeholder="Name of Anime"
          />
         
              
          
                   
          
          <div v-if="v$?.animeName?.$dirty && !v$?.animeName?.$error">The anime invalid</div>
          <div v-if="v$?.animeName?.$dirty && v$?.animeName?.$error">The anime fild is required</div>

          

          <label for="release-date" class="label-inp">Release date</label>
          <input
          v-model.trim = "MyAnime.releaseDate"
          class="imp-create"
          type="number"
          placeholder="When it was released">

          <div v-if="v$?.releaseDate?.$dirty && !v$?.releaseDate?.$error">The anime invalid</div>
          <div v-if="v$?.releaseDate?.$dirty && v$?.releaseDate?.$error">The anime fild is required</div> 

          <label for="quantity-episodios" class="label-inp">Quantity episodios</label>
          <input
          v-model.trim = "MyAnime.quantityEpisodios"
          class="imp-create"
          type="text"
          placeholder="Episodios released">

          <div v-if="v$?.quantityEpisodios?.$dirty && !v$?.quantityEpisodios?.$error">The anime invalid</div>
          <div v-if="v$?.quantityEpisodios?.$dirty && v$?.quantityEpisodios?.$error">The anime fild is required</div>

          <label for="genre-type" class="label-inp">Genre type</label>
          <input
          v-model.trim = "MyAnime.genreType"
          class="imp-create"
          type="text"
          placeholder="Genre">

           <div v-if="v$?.genreType?.$dirty && !v$?.genreType?.$error">The anime invalid</div>
          <div v-if="v$?.genreType?.$dirty && v$?.genreType?.$error">The anime fild is required</div> 

          <label for="written-by" class="label-inp">Written by</label>
          <input
          v-model.trim = "MyAnime.writtenBy"
          class="imp-create"
          type="text"
          placeholder="written by">

           <div v-if="v$?.writtenBy?.$dirty && !v$?.writtenBy?.$error">The anime invalid</div>
          <div v-if="v$?.writtenBy?.$dirty && v$?.writtenBy?.$error">The anime fild is required</div>

          <label for="description" class="label-inp">Description</label>
          <input
          v-model.trim = "MyAnime.description"
          class="imp-create"
          type="text"
          placeholder="Description">

           <div v-if="v$?.description?.$dirty && !v$?.description?.$error">The anime invalid</div>
          <div v-if="v$?.description?.$dirty && v$?.description?.$error">The anime fild is required</div> 
          
          <!-- <span
          class="s-big"
          v-for="error of v$.state.description.$errors"
          :key="error.$uid"
          >{{ description.value }}
          </span> -->          
          
          <button 
            type = "button"           
            id="btn-post"
            v-on:click="sendForm()"            
            value="POST">
            POST
          </button>

        </form>

      </div>

    </main>

  </template>
  
  <style scoped>

main{
  background-color:#344960;
}  

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
  background-color: crimson;
  color: aliceblue;
  
}       

.error-msg{
  color: rgb(236, 51, 51);
  font-size: 13px;
}

@media (min-width:700px){
 .about{
  text-align: left;
  width: 80%;
  margin: auto;
 }

  .imp-create{
  width: 35%;
  margin: initial;
    
}

#btn-post{
  width: 35%;
  margin: initial;


}

}
  
 
  </style>