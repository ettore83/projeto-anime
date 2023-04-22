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
      const response = await axios.get('https://api.example.com/composition')
      state.data = response.data
    } catch (error) {
      state.error = error
    }

    state.loading = false
  }

  return { state, fetchData }
}

export { useCompositionApi }



/*export const APISettings = {
    token: "",
    headers: {
        accept: "applicattion/Json", "X-Api-Key": "5ropeyFvE4d_qg3iWfPaKt8bAmn7cjRQ", 
        //next line comand to help upload image  
        "Content-Type": "multipart/form-data",
     },
      baseURL: "https://api.intern.d-tt.nl/api/",
    }; */