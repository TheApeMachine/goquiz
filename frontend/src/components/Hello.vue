<template>
  <div class="hello">
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
      account: null
    }
  },

  methods: {
    getAccount (id) {
      apiService.getAccount(id).then((data) => {
        this.account = data.data
      })
    }
  },

  mounted () {
    this.getAccount(10000)
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
