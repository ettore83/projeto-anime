import axios from 'axios'
import { reactive } from 'vue'

export default function useCompositionApi() {
  const state = reactive({
    data: null,
    loading: false,
    error: null
  })

  async function fetchData() {
    state.loading = true

    try {
      const response = await axios.get('/api/anime')
      console.log(response)
      state.data = response.data
    } 
      catch (error) {
      console.log(error)
      state.error = error
    }

    state.loading = false
  }

  return { state, fetchData }
}

export { useCompositionApi }


