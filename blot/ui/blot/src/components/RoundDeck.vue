<template>
  <div id="round-deck-table" class="round-deck-table">
    <div class="round-deck">
      <draggable-resizable-vue
          v-for="(card, index) in cards"
          :key="card.id"
          class="card-container"
          v-model:x="card.x"
          v-model:y="card.y"
          v-model:w="card.width"
          :resizable="false"
          parent="#round-deck-table"
          :z="card.zIndex"
      >
        <div :style="card.style">
          <Card :rank="card.rank" :suit="card.suit" />
        </div>
      </draggable-resizable-vue>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import DraggableResizableVue from 'draggable-resizable-vue3';
import Card from './Card.vue';

function getRandomNumber(max: number): number {
  return Math.floor(Math.random() * max);
}

function randomAngle() {
  return Math.floor(Math.random() * 20) - 10;
}

const cards = ref<Array<{
  id: string;
  rank: string;
  suit: string;
  style: {};
  x: number;
  y: number;
  width: number;
  height: number;
  zIndex: number;
}>>([]);

// Здесь вы можете указать свои карты
const cardData = [
  { rank: 'a', suit: 'hearts' },
  { rank: 'k', suit: 'diamonds' },
  { rank: 'q', suit: 'clubs' },
  { rank: 'q', suit: 'spades' },
];

cards.value = cardData.map((card, index) => {
  const id = `${card.rank}-${card.suit}`;
  const angle = randomAngle();
  return {
    id,
    rank: card.rank,
    suit: card.suit,
    style: {
      transform: `rotate(${angle}deg)`,
    },
    // x: getRandomNumber(200),
    // y: getRandomNumber(50),
    x: 400, // TODO find a way to calculate it based on the parent container size. Should be in the middle
    y: 100, // TODO find a way to calculate it based on the parent container size. Should be in the middle
    width: 100, // TODO find a way to calculate it based on the parent container size
    zIndex: index,
  };
});
</script>

<style scoped>
.round-deck-table {
  width: 100%;
}
.round-deck {
  display: flex;
  position: relative;
  width: 100%;
  height: 100vh;
}

.round-deck .card-container {
  width: 60%;
  position: absolute;
  top: 0;
  left: 0;
  height: auto;
  cursor: move;
  transition: transform 0.2s ease;
  border:none;
}

</style>
