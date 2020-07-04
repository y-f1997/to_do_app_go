<template>
<div>
  <div class="container">
    <div class="card-body text-center">
      <div class="form-group">
        <div class="row">
          <div class="col-md-6">
                <label for="startDateTime" class="float-left">開始日時</label>
          </div>
        </div>
        <VueCtkDateTimePicker v-model="postingWhatDone.startDateTime" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'" id="startDateTime"/>
      </div>

      <div class="form-group">
        <div class="row">
          <div class="col-md-6">
                <label for="endDateTime" class="float-left" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'">終了日時</label>
          </div>
        </div>
        <VueCtkDateTimePicker v-model="postingWhatDone.endDateTime" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'" id="endDateTime"/>
      </div>

      <div class="form-group">
        <label for="category" class="float-left">カテゴリ</label>
        <select class="form-control" id="category" placeholder="仕事、食事、etc" v-model="postingWhatDone.category">
          <option value="1">食事</option>
          <option value="2">勉強</option>
          <option value="3">仕事</option>
          <option value="4">遊び</option>
        </select>
      </div>

      <div class="form-group">
        <label for="whatDone" class="float-left">やったこと</label>
        <textarea class="form-control" id="whatDone" v-model="postingWhatDone.whatDone"></textarea>
      </div>

      <div class="form-group">
        <label for="additionalInfo" class="float-left">追加情報</label>
        <textarea class="form-control" id="additionalInfo" v-model="postingWhatDone.additionalInfo"></textarea>
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
      urlPostWhatDone: '/saveWhatDone',
      urlUpdateWhatDone: '/updateWhatDone',
      urlGetWhatDoneDetail: '/getWhatDoneDetail',
      postingWhatDone: {
        userId: null,
        category: null,
        whatDone: null,
        additionalInfo: null,
        startDateTime: null,
        endDateTime: null
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
        this.postingWhatDone.startDateTime = this.$route.params.postStartDateTime
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
      this.$router.push({name:'Home'})
    }else{
      //初期値は更新日時が格納される
      this.keyTime = action;
      this.action='update';
    }
    },
    submit: function () {
      var self = this;
      var req = self.postingWhatDone;

      if(req.startDateTime > req.endDateTime){
        alert("開始日時が終了日時より大きいです");
        return
      }
      console.log(req.startDateTime, req.endDateTime)
      axios.post(self.urlPostWhatDone, req).then(function (res) {
        console.log(res);
        self.$router.push({name:'Home'});
      }).catch(function(err){
        console.err(err);
      })
    },
    update: function(){
      var self = this;
      var req = self.postingWhatDone;
      if(req.startDateTime > req.endDateTime){
        alert("開始日時が終了日時より大きいです");
        return
      }
      console.log(req.startDateTime, req.endDateTime)
      axios.post(self.urlUpdateWhatDone, req).then(function (res) {
        console.log(res);
        self.$router.push({name:'Home'});
      }).catch(function(err){
        console.err(err);
      })


    },
    getDetail: function(){
      var self = this;
      axios.get(this.urlGetWhatDoneDetail, {
        params:{
        whatDoneKey: self.keyTime
      }}).then(function(res){
        self.postingWhatDone = res.data;
      })
    }
  }
}
</script>
