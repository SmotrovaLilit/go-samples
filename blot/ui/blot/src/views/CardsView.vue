<template>
  <div>
    <h1>Your cards</h1>
    <CardDeck :cards="deckOfCards" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import CardDeck from '../components/CardDeck.vue';

export default defineComponent({
  components: {
    CardDeck
  },
  setup() {
    const deckOfCards = [
      { rank: 'k', suit: 'spades' },
      { rank: 'j', suit: 'clubs' },
      { rank: 'q', suit: 'hearts' },
      { rank: 'j', suit: 'diamonds' },
      { rank: 'a', suit: 'spades' },
      { rank: '10', suit: 'clubs' },
      { rank: '7', suit: 'hearts' },
      { rank: '8', suit: 'diamonds' }
    ];

    return {
      deckOfCards
    };
  }
});
</script>
