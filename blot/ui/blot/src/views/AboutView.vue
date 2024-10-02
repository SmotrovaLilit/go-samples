<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div>
      <Card rank="king" suit="hearts" />
      <Card rank="queen" suit="spades" />
      <Card rank="jack" suit="diamonds" />
    </div>
  </div>

</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
<script>
import Card from '../components/Card.vue';

export default {
  components: {
    Card
  }
};
</script>
