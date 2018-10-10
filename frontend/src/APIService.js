import axios from 'axios'
const API_URL = 'http://localhost:6767'

export class APIService {
  getAccount (id) {
    const url = `${API_URL}/accounts/${id}`
    return axios.get(url).then(response => response.data)
  }
}
