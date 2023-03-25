<script setup>
import Demogrid from '../components/TheWelcome.vue'
import { ref } from "vue";

const searchQuery = ref('')
const gridColumns = ['ranking', 'release',]

const gridData = [
  { ranking: 'Ettore', release: 1983 },
  { ranking: 'Priscila', release: 1986 },
  { ranking: 'Lorenzo', release: 2015 },
  { ranking: 'Vicenzo', release: 2019 },
  { ranking: 'luiz', release: 1957 },
  { ranking: 'maria', release: 1959 },
]

const newAnime = ref("")


//fim

/* antigo

const sortOrders = ref(
  orderColumns.value.reduce((o, key) => ((o[key] = 1), o), {})
);

function ageOrder() {
  sortKey.value = 'age';  
}

function rankingOrder(){
  sortKey.value = 'ranking';
}*/

</script>

<template>
  <main id="main-content">
    
    <section id="secSearch">
      <div id="divSearch">
        <input id="inSearch"
         type="text" 
         placeholder="search for anime"
         v-model.trim="searchQuery">
        <div id="bt-div">
          <RouterLink to="/create"><button 
        v-on:click="ageOrder()"
        :class="{ active: sortKey == 'ranking' }"
        id="btRanking"> CREATE NEW        
          <span v-if="sortKey == 'ranking' && sortOrders.ranking < 0">&darr;</span>
          <span v-if="sortKey == 'ranking' && sortOrders.ranking >= 0">&uarr;</span>
          </button></RouterLink>

        <button 
        v-on:click="ageOrder()"
        :class="{ active: sortKey == 'age' }"
          id="btAge" > TO CHOOSE
          <span v-if="sortKey == 'age' && sortOrders.age < 0">&darr;</span>
          <span v-if="sortKey == 'age' && sortOrders.age >= 0">&uarr;</span>
        </button>
      
      </div>    
      
      </div>
      <Demogrid   
          :data="gridData"
          :columns="gridColumns" 
          :filter-key="searchQuery"        
          >
      </Demogrid>
    </section>    
  
  </main>
</template>

<style scoped>

main{
  background-color:var(--color-background-home) ;
  display: grid;
  margin: auto; 
  height: 1000px;
  
}

#secSearch{
  width: 100%;
  margin: auto;
  height: 60px;   
  margin-top: 20px; 
}

#divSearch{
  width: 80%;
  margin: auto;
  margin-bottom: 0px;
  
}

#bt-div{
  box-shadow: 1px 1px 5px rgba(0, 0, 0, 0.406);
}

#inSearch{  
  width: 100%;
  padding-left:25px;       
  height:30px;                                   
  background-image: url('../assets/ic_search@3x.png');  
  background-color: var(--color-text-input);      
  background-size: 12px;
  background-position: left 10px center;
  background-repeat:no-repeat; 
  border-radius: 5px;
  border: none;
  margin: auto;   
  margin-bottom: 10px;
  box-shadow: 1px 1px 5px rgba(0, 0, 0, 0.406);
}

#btRanking{
  width: 50%;
  height:30px;
  border-radius: 5px 0px 0px 5px;
  border: none;
  background-color: var(--color-text-input); 
  color: var(--color-text-button);
  margin: auto;
  cursor: pointer;
}

#btAge{
  width: 50%;
  height:30px;
  border-radius: 0px 5px 5px 0px;
  border: none;
  background-color: var(--color-text-input); 
  color: var(--color-text-button);
  margin: auto;
  cursor: pointer;
}

#bt-div .active{
  background-color: var(--color-button);
}

#div-list{  
  width: 80%;
  height: 110px;
  background-color: #b7d8f4fc;
  border-radius: 10px;
  margin: auto;
  display: flex;
  align-items: center;
  gap: 20px;  
}

img{
  width: 80px;
  height: 100px;
  padding: 5px;
  border-radius: 50%;
}

@media (min-width:700px){

  #secSearch{
    width: 100%;
    margin: auto;
    margin-top: 20px;
  }

  #divSearch{
    width: 80%;
    display: flex;   
    justify-content: space-between;     
  }

  #bt-div{
    display: flex;        
  }

  #inSearch{
    margin: 0;
    width: 35%;
  }

  #btRanking ,#btAge {
    width:150px ;      
  }
  
  
}
</style>
