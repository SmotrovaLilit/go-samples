<template>
  <div class="card">
    <div class="template" :style="templateStyle"></div>
    <div class="rank" :style="rankStyle"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';

export default defineComponent({
  props: {
    rank: {
      type: String,
      required: true
    },
    suit: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const spriteUrl = '/cards/sprite.png';
    const spriteSize = { width: '3584px', height: '6032px' };
    const cardSize = { width: '896px', height: '1302px' };

    const suitPositions: Record<string, { x: number, y: number }> = {
      spades: { x: 0, y: 0 },
      clubs: { x: -896, y: 0 },
      hearts: { x: -1792, y: 0 },
      diamonds: { x: -2688, y: 0 }
    };

    const rankPositions: Record<string, { x: number, y: number }> = {
      king: { x: 0, y: -1302 },
      jack: { x: -896, y: -1302 },
      queen: { x: -1792, y: -1302 }
    };

    const templateStyle = computed(() => ({
      backgroundImage: `url(${spriteUrl})`,
      backgroundPosition: `${suitPositions[props.suit]?.x}px ${suitPositions[props.suit]?.y}px`,
      width: cardSize.width,
      height: cardSize.height,
      backgroundSize: `${spriteSize.width} ${spriteSize.height}`,
      position: 'absolute',
      top: '0',
      left: '0'
    }));

    const rankStyle = computed(() => ({
      backgroundImage: `url(${spriteUrl})`,
      backgroundPosition: `${rankPositions[props.rank]?.x}px ${rankPositions[props.rank]?.y}px`,
      width: cardSize.width,
      height: cardSize.height,
      backgroundSize: `${spriteSize.width} ${spriteSize.height}`,
      position: 'absolute',
      top: '0',
      left: '0'
    }));

    return {
      templateStyle,
      rankStyle
    };
  }
});
</script>

<style scoped>
.card {
  position: relative;
  width: 896px;
  height: 1302px;
  border: 1px solid black;
  border-radius: 8px;
  //background-color: #ccd3d9;
  background-color: #494b4d;
}

.template, .rank {
  background-repeat: no-repeat;
}
</style>
