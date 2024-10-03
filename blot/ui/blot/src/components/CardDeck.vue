<template>
  <div class="deck">
    <div v-for="(card, index) in cardsNew" :key="index" class="card-container"
         :style="card.style">
      <Card :rank="card.rank" :suit="card.suit"/>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import Card from './Card.vue';

function calculateAngle(cardIndex: number, totalCards: number) {
  const angle = 10;
  const half = Math.floor(totalCards / 2);
  const distance = cardIndex - half;
  return distance * angle;
}

function calculateMarginTop(angle: number) {
  return Math.abs(Math.tan(Math.abs(angle) * Math.PI / 180) * 100);
}


export default defineComponent({

  components: {
    Card
  },
  props: {
    cards: {
      type: Array as () => Array<{ rank: string; suit: string }>,
      required: true
    },
    turnOnCards: {
      type: Boolean,
      required: true
    },
    marginBetweenCards: {
      type: Number,
      default: -100,
      required: false
    }
  },
  setup(props) {
    let cardsNew: Array<{ rank: string; suit: string; style: {} }> = [];
    props.cards.forEach((card, index) => {
      let c = {
        rank: card.rank,
        suit: card.suit,
        style: {}
      };
      if (props.turnOnCards) {
        const angle = calculateAngle(index, props.cards.length);
        const marginTop = calculateMarginTop(angle);
        c.style = {
          transform: `rotate(${angle}deg)`,
          marginTop: `${marginTop}px`
        };
      }
      c.style = {
        ...c.style,
        marginRight: `${props.marginBetweenCards}px`
      };
      cardsNew.push(c);
    });

    return {
      cardsNew: cardsNew
    };

  }
});
</script>

<style scoped>
.deck {
  display: flex;
  width: 100%;
  //justify-content: center;
  align-items: center;
  padding-left: 100px;
  padding-right: 200px;
}

.card-container {
  width: 100%;
}
</style>
