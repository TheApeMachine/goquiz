<template>
  <div class="quiz">
    <h1>{{quiz.name}}</h1>
    <p>Question: {{quiz.step}}</p>

    <table class="table table-bordered table-hover">
      <tbody>
        <tr v-for="option in options">
          <td>
            <router-link :to="{name: 'quizzes', params: {id: quiz.id, step: next_step}}">
              {{ option.name }}
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
  name: 'quiz',
  data () {
    return {
      account: null,
      quiz: null,
      step: 0,
      next_step: 0,
      options: []
    }
  },

  methods: {
    getAccount (id) {
      apiService.getAccount(id).then((data) => {
        this.account = data.data
      })
    },

    getQuiz (id) {
      apiService.getQuiz(id).then((data) => {
        this.quiz = data
        this.options = data.Steps[this.$route.params.step].Options
      })
    }
  },

  mounted () {
    this.getAccount(10000)
    this.getQuiz(this.$route.params.id)
    this.step = this.$route.params.step
    this.next_step = this.$route.params.step + 1
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
