<template>
  <div class="game-container">
    <div class="game-top">
    </div>
    <div class="game-middle">
      <div class="left-bar"></div>
      <div class="middle">
<!--        <div class="table-cards"><CardDeck :cards="cardsInRound" :turn-on-cards=false :margin-between-cards=-100 /></div>-->
      </div>
      <div class="right-bar"></div>
    </div>
    <div class="game-bottom">
      <div class="your-cards">
        <CardDeck :cards="myCards" :turnOnCards=true />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import CardDeck from '../components/HandDeck.vue';

export default defineComponent({
  components: {
    CardDeck
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
  //background-color: gray;
}
.table-cards {
  width: 60%;
  display: flex;
  justify-content: center;
}
.your-cards {
  width: 60%;
  display: flex;
  justify-content: center;
}
</style>