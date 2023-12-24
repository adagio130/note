# Declarative Rendering
- using a template syntax that extends HTML, we can describe how the HTML should look like based on JavaScript state.
- `reative()` only works on objects (including arrays and built-in types like `Map` and `Set`
- `ref()` can take any value type and create an object that exposes the inner value under a `.value` property

# Attribute Bindings
- `v-bind` 
- `<div :id="dynamicId"></div>`

# Event Listeners
- `v-on`
- `<button @click="increment">{{ count }}</button>`

# Form Bindings
- `<input :value="text" @input="onInput">`
- two-way bindings  `<input v-model="text">` 

# Conditional Rendering
- `<h1 v-if="awesome">Vue is awesome!</h1>` `
- `<h1 v-else>Oh no 😢</h1>`

