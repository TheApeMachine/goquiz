<template>
  <div class="hello">
    <table class="table table-bordered table-hover">
      <tbody>
        <tr v-for="quiz in quizzes">
          <td>
            <router-link :to="{name: 'quizzes', params: {id: quiz.id, step: 0}}">
              {{ quiz.name }}
            </router-link>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import {APIService} from '../APIService'
const apiService = new APIService()

export default {
  name: 'hello',
  data () {
    return {
      msg: 'Welcome to Your Vue.js PWA',
      account: null,
      quizzes: []
    }
  },

  methods: {
    getAccount (id) {
      apiService.getAccount(id).then((data) => {
        this.account = data.data
      })
    },

    getQuizzes () {
      apiService.getQuizzes().then((data) => {
        this.quizzes = data
        console.log(data)
      })
    }
  },

  mounted () {
    this.getAccount(10000)
    this.getQuizzes()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #35495E;
}
</style>
