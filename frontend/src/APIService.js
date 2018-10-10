import axios from 'axios'
const API_URL = 'http://localhost'
const ACCOUNT_PORT = 6767
const QUIZ_PORT = 6768

export class APIService {
  getAccount (id) {
    const url = `${API_URL}:${ACCOUNT_PORT}/accounts/${id}`
    return axios.get(url).then(response => response.data)
  }

  getQuizzes () {
    const url = `${API_URL}:${QUIZ_PORT}/quizzes`
    return axios.get(url).then(response => response.data)
  }
}
