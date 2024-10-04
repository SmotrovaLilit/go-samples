<template>
  <div class="players-cards top-player-cards">
    <ClosedDeck :cards-count="cardsInAHand"/>
  </div>
  <div class="players-cards left-player-cards">
    <ClosedDeck :cards-count="cardsInAHand"/>
  </div>
  <div class="players-cards right-player-cards">
    <ClosedDeck :cards-count="cardsInAHand"/>
  </div>
  <div class="game-container">
    <div class="game-top">

    </div>
    <div class="game-middle">
      <div class="left-bar">
      </div>
      <div class="middle">
        <div class="game-table">
          <div class="table-cards">
            <RoundDeck :cards="cardsInRound"/>
          </div>
        </div>
      </div>
      <div class="right-bar">
      </div>
    </div>
    <div class="game-bottom">
      <div class="your-cards">
        <HandDeck :cards="myCards"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import HandDeck from "@/components/HandDeck.vue";
import RoundDeck from "@/components/RoundDeck.vue";
import ClosedCard from "@/components/ClosedCard.vue";
import ClosedDeck from "@/components/ClosedDeck.vue";

export default defineComponent({
  components: {
    ClosedDeck,
    ClosedCard,
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
      {rank: '8', suit: 'diamonds'}
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
      cardsInAHand: 8
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
  display: flex;
  justify-content: center;
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
  //display: flex;
  //justify-content: center;
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
  display: flex;
  justify-content: center;
  //background-color: green;
}

.game-bottom {
  display: flex;
  //position: fixed;
  //bottom: 0;
  //left: 0;
  //width: 100%;
  //height: 30%;
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
  position: fixed;
  bottom: 0;
  padding-bottom: 60px;
  left: 50%;
  transform: translateX(-50%);
}

.players-cards {
  width: 20%;
  position: fixed;
}

.left-player-cards {
  left: 0;
  top: 50%;
  transform: translateY(-50%) rotate(90deg);
}


.right-player-cards {
  right: 0;
  top: 50%;
  transform: translateY(-50%) rotate(-90deg);
}

.top-player-cards {
  top: 0;
  left: 50%;
  transform: translateX(-50%) rotate(-180deg) ;
}
</style>