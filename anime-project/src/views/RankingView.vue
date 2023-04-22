<script>
import { useCompositionApi } from '../services/config'

export default {
  setup() {
    const compositionApi = useCompositionApi()

    compositionApi.fetchData()

    return { compositionApi }
  }
}
</script>

<template>
    <main>
      <div>
        <h1>Avalie o anime</h1>
        <div id="stars">
        
        </div>
      </div>    
      <p v-if="compositionApi.state.loading">Carregando...</p>
    <p v-if="compositionApi.state.error">Erro ao carregar os dados</p>
    <ul v-if="compositionApi.state.data">
      <li v-for="item in compositionApi.state.data" :key="item.id">{{ item.name }}</li>
    </ul> 

      <div>
      
      </div>

    </main>
  </template>
  
  <style scoped>
  .about{
    width: 80%;    
    margin: auto;    
  }

  #stars{
    width: 80%;
    height: 80px;
    border: 2px solid black;
  }
  
  </style>