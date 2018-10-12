<template>
  <div class="quiz">
    {{quiz.name}}

    <table class="table table-bordered table-hover">
      <tbody>
        <tr v-for="option in options">
          <td>
             {{ option.name }}
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
        this.options = this.quiz.Options
      })
    }
  },

  mounted () {
    this.getAccount(10000)
    this.getQuiz(this.$route.params.id)
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
