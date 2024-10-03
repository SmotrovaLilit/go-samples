<template>
  <div class="game-container">
    <div class="game-top">
    </div>
    <div class="game-middle">
      <div class="left-bar"></div>
      <div class="middle">
        <div class="game-table">
          <div class="table-cards"><RoundDeck :cards="cardsInRound" /></div>
        </div>
      </div>
      <div class="right-bar"></div>
    </div>
    <div class="game-bottom">
      <div class="your-cards">
        <HandDeck :cards="myCards" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import HandDeck from "@/components/HandDeck.vue";
import RoundDeck from "@/components/RoundDeck.vue";

export default defineComponent({
  components: {
    RoundDeck,
    HandDeck,
  },
  setup() {
    const myCards = [
      {rank: 'k', suit: 'spades'},
      {rank: 'j', suit: 'clubs'},
      {rank: 'q', suit: 'hearts'},
      {rank: 'j', suit: 'diamonds'},
      {rank: 'a', suit: 'spades'},
      {rank: '10', suit: 'clubs'},
      {rank: '7', suit: 'hearts'},
      { rank: '8', suit: 'diamonds' }
    ];
    const cardsInRound = [
      {rank: 'k', suit: 'spades'},
      {rank: 'j', suit: 'clubs'},
      {rank: 'q', suit: 'hearts'},
      {rank: 'j', suit: 'diamonds'},
      // {rank: 'a', suit: 'spades'},
      // {rank: '10', suit: 'clubs'},
      // {rank: '7', suit: 'hearts'},
      // { rank: '8', suit: 'diamonds' }
    ];

    return {
      cardsInRound,
      myCards,
    };
  }
});
</script>
<style scoped>
.game-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.game-top {
  height: 20%;
  //background-color: blue;
}

.game-middle {
  display: flex;
  flex-direction: row;
  height: 50%;
  //background-color: red;
}
.left-bar {
  width: 20%;
  //background-color: white;
}

.middle {
  width: 60%;
  display: flex;
  justify-content: center;
  padding: 20px;
  //background-color: yellow;
}

.right-bar {
  width: 20%;
  //background-color: green;
}

.game-bottom {
  display: flex;
  //position: fixed;
  //bottom: 0;
  //left: 0;
  //width: 100%;
  //height: 30%;
  justify-content: center;
  padding-bottom: 30px;
  box-sizing: border-box;
  //background-color: gray;
}
.game-table {
  display: flex;
  justify-content: center;
  width: 100%;
  background-color: #4f0c09;
  border-radius: 40%;
  box-shadow: 0 0 20px 10px #4f0c09;

}

.table-cards {
  width: 100%;
  display: flex;
  justify-content: center;
}
.your-cards {
  width: 40%;
  display: flex;
  justify-content: center;
  position: fixed;
  bottom: 0;
  padding-bottom: 60px;
}
</style>