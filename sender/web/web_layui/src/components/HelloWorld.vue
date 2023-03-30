<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h1>{{ name }}</h1>
  </div>
  <lay-button type="primary">原始按钮</lay-button>
  <lay-button>默认按钮</lay-button>
  <lay-button type="normal">百搭按钮</lay-button>
  <lay-button type="warm">暖色按钮</lay-button>
  <lay-button type="danger">警告按钮</lay-button>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String,
    name: String
  }
}
</script>

