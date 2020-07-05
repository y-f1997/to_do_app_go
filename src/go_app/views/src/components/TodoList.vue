<template>
    <div class="container-fluid">
    <div class="row">
      <div class="col-md-11">
      </div>
      <div class="col-md-1">
      <h2 class="float-right">
        <div @click="routeDetail(null)" style="color:blue; cursor:pointer">
          <font-awesome-icon icon="file-alt" />
        </div>
      </h2>

      </div>
    </div>
    <div class="row">
      <TodoTable :todo_list="listTodo" todo_status="todo" @parent-get-list="getList"></TodoTable>
    </div>
    <div class="row">
      <TodoTable :todo_list="listFinished" todo_status="fin" @parent-get-list="getList"></TodoTable>
    </div>

  </div>

</template>

<script>
import axios from 'axios';
import TodoTable from './todo/TodoTable.vue'
export default {
  name: 'Home',
  components:{
    TodoTable
  },
  data() {
    return {
      urlGetList: "/getTodoList",
      list:[],
      listTodo: [],
      listFinished: [],
      selectedDate: {
        start:new Date(),
        end:new Date()
        }
    }
  },
  mounted: function () {
    this.getList()
  },

  methods: {
    getList() {
      var self = this;
      axios.get(this.urlGetList).then(function (res) {
        var list = res.data;
        if (list === null) {
          list = [];
        }
        self.listTodo = list.filter(v => {
          return v.timeDone === null;
        });
        self.listFinished = list.filter(v => {
          return v.timeDone !== null;
        });
      })
    },
    routeDetail: function (todo) {
      var lastWhatDone = {};
      if (this.list.length>0){
        lastWhatDone = this.list[this.list.length - 1];
      }else{
        lastWhatDone.endDateTime = null;
      }
      
      var to = {
        //新規作成用URL
        name: 'Todo',
        params: {
          action: 'post',
          postStartDateTime: lastWhatDone.endDateTime
        }
      }
      if (todo !== null) {
        //更新用URL(crtTimestampでサーチ)
        to.params.action = todo.crtTimestamp;
      }
      this.$router.push(to)
    },
  }

}
</script>

<style>
  table{
    margin-top: 1em;
  }

  .list {
    table-layout: fixed;
  }
</style>