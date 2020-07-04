<template>
<div>
  <div id="error" v-html="error" v-if="error"></div>

  <div class="row" style="margin-top: 200px;">
    <div class="card mx-auto" style="width:400px;">
      <div class="card-header">Sign Up</div>
      <div class="card-body text-center">
        <div class="form-group">
          <label for="inputEmail">User ID</label>
          <input type="email" class="form-control" id="inputEmail" aria-describedby="emailHelp" placeholder="Enter User ID" v-model="userId">
          <label for="inputPassword">Password</label>
          <input type="password" class="form-control" id="inputPassword" placeholder="Password" v-model="password">
          <label for="name">氏名</label>
          <input type="text" class="form-control" id="name" placeholder="name" v-model="name">
        </div>
        <button class="btn btn-primary mx-auto" @click="signUp()">SignUp</button>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SignUp',
  data() {
    return {
      error: null,
      urlSignUp: "/signUp",
      userId: null,
      password: null,
      name: null
    }
  },
  methods: {
    signUp: function () {
      var self = this;
      var req = {
        userId: this.userId,
        password: this.password,
        name: this.name
      }
      axios.post(this.urlSignUp, req).then(function (res) {
        console.log(res);
        self.$router.push({
          name: 'Login'
        })
      }).catch(error => {
        self.error = "<div class='alert alert-warning'>" + error.response.data.code + " :" + error.response.data.error + "</div>"
      })
    }
  }
}
</script>
