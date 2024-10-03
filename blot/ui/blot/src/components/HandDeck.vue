<template>
  <draggable
      :list="cards"
      :disabled="false"
      item-key="id"
      class="deck"
      ghost-class="ghost"
      :move="checkMove"
      @end="onDragEnd"
      @start="onDragStart"
  >
    <template #item="{ element }">
      <div class="card-container" :class="{ 'not-draggable': false }"
           :style="element.style">
        <Card :rank="element.rank" :suit="element.suit"/>
      </div>
    </template>
  </draggable>
</template>

<script lang="ts">
import {defineComponent, ref, watch} from 'vue';
import Card from './Card.vue';
import draggable from "vuedraggable";

function calculateAngle(cardIndex: number, totalCards: number) {
  const angle = 10;
  const half = Math.floor(totalCards / 2);
  const distance = cardIndex - half;
  return distance * angle;
}

function calculateMarginTop(angle: number) {
  return Math.abs(Math.tan(angle * Math.PI / 180) * 100);
}

function calculateNewPositions(
    length: number,
    currentIndex: number,
    futureIndex: number
): number[] {
  if (currentIndex < 0 || futureIndex < 0 || currentIndex >= length || futureIndex >= length) {
    return array;
  }
  let res: number[] = new Array(length).fill(0);
  for (let i = 0; i < length; i++) {
    if (i === currentIndex) {
      res[currentIndex] = futureIndex;
    } else if (currentIndex < futureIndex) { // moving right
      if (i > currentIndex && i <= futureIndex) { // changed range
        res[i] = i - 1;
      } else { // no changed range
        res[i] = i;
      }
    } else { // moving left
      if (i >= futureIndex && i < currentIndex) { // changed range
        res[i] = i + 1;
      } else {
        res[i] = i  // no changed range
      }
    }
  }
  return res;

}

export default defineComponent({
  components: {
    draggable,
    Card
  },
  props: {
    cards: {
      type: Array as () => Array<{ rank: string; suit: string }>,
      required: true
    },
    marginBetweenCards: {
      type: Number,
      default: -100,
      required: false
    }
  },
  setup(props) {
    const cards = ref<Array<{
      id: string;
      rank: string;
      suit: string;
      style: {}
    }>>([]);
    const rotateCards = () => {
      cards.value.forEach((card, index) => {
        const angle = calculateAngle(index, props.cards.length);
        const marginTop = calculateMarginTop(angle);
        card.style.transform = `rotate(${angle}deg)`;
        card.style.marginTop = `${marginTop}px`;
        card.style.zIndex = index;
      });
    };
    cards.value = props.cards.map((card, index) => {
      const id = `${card.rank}-${card.suit}`;
      return {
        id,
        rank: card.rank,
        suit: card.suit,
        style: {
          transform: ``,
          marginTop: ``,
          marginRight: `${props.marginBetweenCards}px`,
          zIndex: 0,
        },
      };
    });
    rotateCards();
    watch(
        () => cards,
        (newValue, oldValue) => {
          console.log('cards changed', newValue, oldValue);
        },
        {deep: true}
    )
    const onDragEnd = (evt) => {
      console.log('ended!', evt);
      rotateCards();
    };

    const checkMove = (e) => {
      console.log('moving', e);
      if (e.draggedContext.futureIndex === undefined) {
        return;
      }
      console.log('Future index: ' + e.draggedContext.futureIndex);
      console.log('index: ' + e.draggedContext.index);

      console.log('Old array: ' + cards.value[0].rank + ' ' + cards.value[0].style.zIndex)
      console.log('Old array: ' + cards.value[1].rank + ' ' + cards.value[1].style.zIndex)
      console.log('old array: ' + cards.value[2].rank + ' ' + cards.value[2].style.zIndex)

      const newPositions = calculateNewPositions(cards.value.length, e.draggedContext.index, e.draggedContext.futureIndex);
      for (let i = 0; i < newPositions.length; i++) {
        const angle = calculateAngle(newPositions[i], cards.value.length);
        const marginTop = e.draggedContext.index == i ? -100 + calculateMarginTop(angle): calculateMarginTop(angle);
        cards.value[i].style.zIndex = newPositions[i];
        cards.value[i].style.transform = `rotate(${angle}deg)`;
        cards.value[i].style.marginTop = `${marginTop}px`;
      }
      console.log(cards.value[0].rank + ' ' + cards.value[0].style.zIndex)
      console.log(cards.value[1].rank + ' ' + cards.value[1].style.zIndex)
      console.log(cards.value[2].rank + ' ' + cards.value[2].style.zIndex)
    };
    const onDragStart = (e) => {
      console.log('started!', e);
      const draggedElement = e.item;
      draggedElement.style.zIndex = cards.value[e.oldIndex].style.zIndex;
      draggedElement.style.transform = cards.value[e.oldIndex].style.transform;
      const angle = calculateAngle(e.oldIndex, cards.value.length);
      const marginTop = -100 + calculateMarginTop(angle);
      cards.value[e.oldIndex].style.marginTop = `${marginTop}px`;
    };
    return {
      cards: cards,
      dragging: true,
      onDragEnd,
      checkMove,
      onDragStart
    };
  }
});
</script>

<style scoped>
.ghost {
  opacity: 0.5;
}

.deck {
  display: flex;
  width: 100%;
  align-items: center;
  padding-left: 100px;
  padding-right: 200px;
}

.card-container {
  width: 100%;
  cursor: move;
  transition: transform 0.2s ease;
}

.not-draggable {
  cursor: no-drop;
}
</style>
