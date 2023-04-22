/* import { APISettings } from "./config";import axios from "axios";

export default {
    index() { console.log(APISettings);
    return axios.get(APISettings.baseURL + "houses", {
        headers: APISettings.headers,
    });
},
    create(params) {console.log(APISettings);
    return axios.post(APISettings.baseURL + "houses", params, {
        headers: APISettings.headers,
    });
},
}
*/