<template>
<div>
  <div class="row" style="margin-top: 200px;">
    <div class="card mx-auto" style="width:400px;">
      <div class="card-body text-center">
        <div class="form-group">
          <label for="inputEmail">User ID</label>
          <input type="email" class="form-control" id="inputEmail" aria-describedby="emailHelp" placeholder="Enter User ID"
          v-model="userId">
          <label for="inputPassword">Password</label>
          <input type="password" class="form-control" id="inputPassword" placeholder="Password"
          v-model="password">
        </div>
        <button class="btn btn-primary mx-auto" @click="login()">Login</button>
      </div>

    </div>

  </div>

</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Login',
  data(){
    return {
      urlIsSession: "/checkIsOnSession",
      urlLogin: "/login", 
      userId: null,
      password: null,
      error: null
    }
  },
  created: function(){
      var self = this;
      axios.get(this.urlIsSession).then(res => {
        if(res.data.onSession === true){
          self.$router.push({name:"Home"})
        }
      })
  },
  methods:{
    login(){
      var req = {
        "userId": this.userId,
        "password": this.password
      }
      var self = this;
      axios.post(this.urlLogin,req).then(res => {
        self.error= null;
        console.log(res)
        self.$router.push("/home")
        })
    },
  }
}
</script>

