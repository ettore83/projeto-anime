<script setup>
//tudo novo
import { ref, computed } from 'vue'

const props = defineProps({
  data: Array,
  columns: Array,
  filterKey: String
})

const sortKey = ref('')
const sortOrders = ref(
  props.columns.reduce((o, key) => ((o[key] = 1), o), {})
)

const filteredData = computed(() => {
  let { data, filterKey } = props
  if (filterKey) {
    filterKey = filterKey.toLowerCase()
    data = data.filter((row) => {
      return Object.keys(row).some((key) => {
        return String(row[key]).toLowerCase().indexOf(filterKey) > -1
      })
    })
  }
  const key = sortKey.value
  if (key) {
    const order = sortOrders.value[key]
    data = data.slice().sort((a, b) => {
      a = a[key]
      b = b[key]
      return (a === b ? 0 : a > b ? 1 : -1) * order
    })
  }
  return data
})

function sortBy(key) {
  sortKey.value = key
  sortOrders.value[key] *= -1  
}

function capitalize(str) {
  return str.charAt(0).toUpperCase() + str.slice(1)
}
</script>

<template>

<main id="main-list">

  <section id="sec-search" v-if = "filteredData.length">
  
    <div id  = "div-thead">

      <div id = "div-tr">  
      
        <button id = "bt-div" 
          v-for = "key in columns"
          @click = "sortBy(key)"
          :class = "{ active: sortKey == key }">

          {{ capitalize(key) }}

          <span class = "arrow" :class="sortOrders[key] > 0 ? 'asc' : 'dsc'"> </span>
        
        </button>
      
      </div>

    </div>
  
    <div id = "div-tbody">
  
      <div id = "div-info" v-for = "entry in filteredData">

        <img src = "../assets/chainsaw-man.jpg" alt = "">
        <div id = "info-name">

          <h1>name</h1>
          <p>{{entry["age"]}}</p>
          <p>episodios</p>
  
        </div>
  
        <div id = "div-td" v-for = "key in columns">

          {{entry[key]}}  
  
        </div>
  
      </div>
  
    </div>

  </section>  
  <p v-else>No matches found.</p>
  
</main>

</template>

<style scoped>

.arrow {
  display: inline-block;
  vertical-align: middle;
  width: 0;
  height: 0;
  margin-left: 5px;
 
}
.arrow.asc {
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-bottom: 4px solid #fff;
}

.arrow.dsc {
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-top: 4px solid #fff;
}

#main-list{
  margin-top: 0px;
  
}

#div-tr{
  width: 80%;
  margin: auto;
  display: flex;
  justify-content: space-between;
}

#bt-div {
  width: 50%;
  height:30px;
  border-radius: 5px 5px 5px 5px;
  border: none;
  background-color: var(--color-button); 
  color: var(--color-text-button);
  margin: auto;
  cursor: pointer;
  margin-top: 25px;
  box-shadow: 1px 1px 5px rgba(0, 0, 0, 0.406);
  margin-left: 0;
  
}
#bt-div {
  margin-right: 0;
}

#div-tbody{
  width: 100%;
  
}


#div-info{  
  width: 80%;
  height: 110px;
  background-color: #3277b3fc;
  border-radius: 10px;
  margin: auto;
  display: flex;
  align-items: center;
  gap: 2px;  
  margin-top: 20px;
}

img{
  width: 80px;
  height: 100px;
  padding: 5px;
  border-radius: 50%;
}


@media (min-width:700px){

#bt-div{
  width: 30%;

  }

}


</style>
