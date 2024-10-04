<template>
  <div id="round-deck-table" class="round-deck-table">
    <div class="round-deck">
      {{ title }}
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
          <Card :rank="card.rank" :suit="card.suit"/>
        </div>
      </draggable-resizable-vue>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineProps, ref} from 'vue';
import DraggableResizableVue from 'draggable-resizable-vue3';
import Card from './Card.vue';

interface Card {
  rank: string;
  suit: string;
}

interface Props {
  cards: { type: Card[], required: false };
}
const props = withDefaults(defineProps<Props>(), {
  cards: [],
})

function generateRandomNumberBetween(min: number, max: number): number {
  console.log('min', min);
  console.log('max', max);
  console.log(Math.floor(Math.random() * (max - min + 1)) + min);
  return Math.floor(Math.random() * (max - min + 1)) + min;
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

const parentWidth = 600; // TODO get the parent container width
const parentHeight = 600; // TODO get the parent container height

cards.value = props.cards.map((card, index) => {
  const id = `${card.rank}-${card.suit}`;
  const angle = randomAngle();

  const cardWidth = 0.25 * parentWidth; // 20% of the parent container width
  const cardX = (parentWidth - cardWidth) / 2; // Center horizontally
  const cardY = (parentHeight - cardWidth) / 2; // Center vertically
  const tolerance = 0.25 * parentWidth; // 5% of the parent container width

  return {
    id,
    rank: card.rank,
    suit: card.suit,
    style: {
      transform: `rotate(${angle}deg)`,
    },
    x: generateRandomNumberBetween(cardX - tolerance, cardX + tolerance), // Centered X position;
    y: generateRandomNumberBetween(cardY - tolerance, cardY + tolerance / 2), // Centered Y position
    width: cardWidth,
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
  //transition: transform 1s ease;
  border: none;
}

</style>
