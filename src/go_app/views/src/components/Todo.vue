<template>
<div>
  <div class="container">
    <div class="card-body text-center">
      <div class="form-group">
        <div class="row">
          <div class="col-md-6">
                <label for="timeToBeDone" class="float-left">期限</label>
          </div>
        </div>
        <VueCtkDateTimePicker v-model="postingTodo.timeToBeDone" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'" id="timeToBeDone"/>
      </div>

      <div class="form-group">
        <div class="row">
          <div class="col-md-6">
                <label for="timeDone" class="float-left" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'">終了日時</label>
          </div>
        </div>
        <VueCtkDateTimePicker v-model="postingTodo.timeDone" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'" id="timeDone"/>
      </div>

      <div class="form-group">
        <label for="importance" class="float-left">重要度</label>
        <select class="form-control" id="importance" placeholder="仕事、食事、etc" v-model="postingTodo.importance">
          <option value="1">最重要</option>
          <option value="2">重要</option>
          <option value="3">不急</option>
          <option value="4">余力</option>
        </select>
      </div>

      <div class="form-group">
        <label for="category" class="float-left">カテゴリ</label>
        <select class="form-control" id="category" placeholder="仕事、食事、etc" v-model="postingTodo.category">
          <option value="1">仕事</option>
          <option value="2">勉強</option>
          <option value="3">生活</option>
          <option value="4">目標</option>
        </select>
      </div>

      <div class="form-group">
        <label for="toDo" class="float-left">ToDo</label>
        <textarea class="form-control" id="toDo" v-model="postingTodo.toDo"></textarea>
      </div>

      <button v-if="action==='post'" class="btn btn-primary mx-auto" @click="submit()">登録</button>
      <button v-if="action==='update'" class="btn btn-primary mx-auto" @click="update()">更新</button>
    </div>
  </div>
</div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'WhatDone',
  data() {
    return {
      urlPostTodo: '/saveTodo',
      urlUpdateTodo: '/updateTodo',
      urlGetTodoDetail: '/getTodo',
      postingTodo: {
        userId: null,
        category: null,
        importance: null,
        toDo: null,
        timeToBeDone: null,
        timeDone: null
      },
      action: null,
      keyTime: null
    }
  },
  created: function(){
    this.switchAction()
    if (this.action === 'update') {
      this.getDetail()
    }
    if (this.action === 'post'){
        this.postingTodo.timeToBeDone = this.$route.params.postStartDateTime
    }

  },
  watch: {
    $route: function(){
      this.switchAction()
    }
  },
  methods: {
    switchAction: function(){
    var action = this.$route.params.action;
    if (action === 'post'){
      this.action='post'
    }else if(!action){
      this.$router.push({name:'TodoList'})
    }else{
      //初期値は更新日時が格納される
      this.keyTime = action;
      this.action='update';
    }
    },
    submit: function () {
      var self = this;
      var req = self.postingTodo;

      axios.post(self.urlPostTodo, req).then(function (res) {
        console.log(res);
        self.$router.push({name:'TodoList'});
      }).catch(function(err){
        console.err(err);
      })
    },
    update: function(){
      var self = this;
      var req = self.postingTodo;

      axios.post(self.urlUpdateTodo, req).then(function (res) {
        console.log(res);
        self.$router.push({name:'TodoList'});
      }).catch(function(err){
        console.err(err);
      })


    },
    getDetail: function(){
      var self = this;
      axios.get(this.urlGetTodoDetail, {
        params:{
        keyTime: self.keyTime
      }}).then(function(res){
        self.postingTodo = res.data;
      })
    }
  }
}
</script>
