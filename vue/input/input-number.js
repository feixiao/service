function isValueNumber (value) {
  return true  // 这边偷懒了
}

Vue.component('input-number',{
  template: '\
    <div class="input-number"> \
      <input \
        type="text" \
        :value="currentValue" \
        @change="handleChange"> \
      <button \
        @click="handleDown" \
        :disabled="currentValue <= min">-</button> \
      <button \
        @click="handleUp" \
        :disabled="currentValue >= max">+</button> \
    </div>',
  props: {
    max: {
      type: Number,
      default: Infinity
    },
    min: {
      type: Number,
      default: -Infinity
    },
    value: {
      type: Number,
      default: 0
    }
  },
  data: function () {
    var val = this.value
    if (val > this.max ) val = this.max;
    if (val < this.min ) val = this.min;

    return {
      currentValue: val
    }
  },
  watch: {
    currentValue: function (val) {
      this.$emit('input',val)
      this.$emit('on-change',val)
    },
    value: function (val) {
      this.updateValue(val)
    }
  },
  methods: {
    handleDown: function() {
      if (this.currentValue <= this.min ) return ;
      this.currentValue -= 1;
    },
    handleUp: function() {
      if (this.currentValue >= this.max) return ;
      this.currentValue += 1;
    },
    updateValue: function (val) {
      if (val > this.max ) val = this.max
      if (val < this.min ) val = this.min
      this.currentValue = val
    },
    handleChange: function (event) {
      var val = event.target.value.trim();
      var max = this.max
      var min = this.min
      if (isValueNumber(val)) {
        val = Number(val)
        this.currentValue = val
        if (val > max ) {
          this.currentValue = max;
        }else if (val < min) {
          this.currentValue = min 
        }
      }else {
        event.target.value = this.currentValue
      }
    }
  },
  mounted: function () {
    this.updateValue(this.value)
  },
});