<template>
  <div class="deck">
    <div v-for="(card, index) in cards" :key="index" class="card-container">
      <Card :rank="card.rank" :suit="card.suit" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Card from './Card.vue';

function calculateAngle(cardIndex: number, totalCards: number) {
  const angle = 10;
  const half = Math.floor(totalCards / 2);
  const distance = cardIndex - half;
  return distance * angle;
}

export default defineComponent({

  components: {
    Card
  },
  props: {
    cards: {
      type: Array as () => Array<{ rank: string; suit: string }>,
      required: true
    }
  },
  setup(props) {
    cards: Array<{ rank: string; suit: string; styles:  }>;
    return {
      cards
    };
  }
});
</script>

<style scoped>
.deck {
  display: flex;
  width: 100%;
}

.card-container {
  width: 100%;
  margin-right: 10px;
}
</style>
