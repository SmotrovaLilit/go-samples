<template>
  <div class="list-container">
      <draggable
          :list="list"
          :disabled="!enabled"
          item-key="name"
          class="list-group"
          ghost-class="ghost"
          :move="checkMove"
          @start="dragging = true"
          @end="dragging = false"
      >
        <template #item="{ element }">
          <div class="list-group-item" :class="{ 'not-draggable': !enabled }">
            <Card :rank="element.rank" :suit="element.suit"/>

          </div>
        </template>
      </draggable>
  </div>
</template>

<script>
import draggable from "vuedraggable";
import Card from "@/components/Card.vue";
let id = 1;
export default {
  name: "simple",
  display: "Simple",
  order: 0,
  components: {
    Card,
    draggable
  },
  data() {
    return {
      enabled: true,
      list: [
        { rank: "q", suit: "diamonds", id: 0 },
        { rank: "j", suit: "clubs", id: 1 },
        { rank: "k", suit: "hearts", id: 2 }
      ],
      dragging: true
    };
  },
  computed: {
  },
  methods: {
    // add: function() {
    //   this.list.push({ name: "Juan " + id, id: id++ });
    // },
    // replace: function() {
    //   this.list = [{ name: "Edgard", id: id++ }];
    // },
    checkMove: function(e) {
      window.console.log("Future index: " + e.draggedContext.futureIndex);
    }
  }
};
</script>
<style scoped>
.buttons {
  margin-top: 35px;
}

.ghost {
  opacity: 0.5;
  background: #c8ebfb;
}

.not-draggable {
  cursor: no-drop;
}
.list-group {
  display: flex;
  flex-direction: row;
  padding: 0;
  margin: 0;
  list-style-type: none;
  border: 1px solid #ccc;
  border-radius: 4px;
  overflow: hidden;
}
.list-group-item {

}
</style>