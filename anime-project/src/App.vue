<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { ref } from "vue";

let menuactive = ref(false);

function openMenu(){
  menuactive.value = true;
}
function closeMenu(){
  menuactive.value = false;
}

</script>

<template>
  <header>   
    
      <nav class="container">
        <img src="../src/assets/logo.png" alt="">

        <img v-on:click="openMenu()" src="./assets/menu-mobile.png" alt="menu" id="menu-button">

       
        <div v-on:click="closeMenu()" id="menu-overlay" v-if="menuactive"></div>

        <div id="menu-itens" :class="{active: menuactive}">

          <img id="menu-logo" src="../src/assets/logo.png" alt="">

          <ul>
              <li><RouterLink to="/">Home</RouterLink></li>
              <li><RouterLink to="/ranking">Ranking</RouterLink></li>
              <li><RouterLink to="/create">Create</RouterLink></li>
              <li><RouterLink to="/about">About</RouterLink></li>          
          </ul>
        
        </div>
      </nav>
    
  </header>

  <RouterView />
</template>

<style scoped>

header{
  background-color:var(--color-background-nav);
  height: 60px;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: auto;
}
nav{
  width: 80%;
  display: flex;
  justify-content: space-between;
  height: 60px;
  align-items: center ;

}
nav img{
  width: 130px;
}

#menu-button{
  width: 40px;
  cursor: pointer;
}

#menu-overlay{
  position: fixed;
  top: 0;
  left: 0;
  width: 40%;
  height: 100vh;
  background-color: black;
  opacity: 0.8;
  cursor: pointer;
}

#menu-logo{
  width: 100px;
  margin-top: 30px;
  margin-bottom: 10px;  
}

#menu-itens{
  position: fixed;
  top: 0;
  right: 0;
  background-color: var(--color-background-nav);
  width: 60%;
  height: 100vh;
  display: none;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  
}

#menu-itens.active{
  display: flex;

}

ul{
  list-style: none;
  text-align: center;
}

ul li{
  margin: 20px 0px;
}

ul li a{
  color: var(--color-text-light);
}

@media (min-width:700px){
  #menu-button,
  #menu-logo,
  #menu-overlay{
    display: none;
  }

  #menu-itens{
    display: flex;
    position: static;
    height: 60px;
    width: auto;
  }

  ul{
    display: flex;  
    height: 60px;  
    align-items: center;
  }

  ul li{
    margin: 0;
    margin-left: 20px;    
  }
}

</style>
