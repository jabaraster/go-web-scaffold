
const FormInputMixin = {
    getInitialState: function() {
        return {
            errors: {}
        };
    },
    hasError: function() {
        for (let p in this.state.errors) {
            if (this.state.errors[p]) {
                return true;
            }
        }
        return false;
    },
    handleValueChange: function(e) {
        this.state.errors[e.descriptor] = e.error;
        if (e.error) {
            return;
        }
        const newValue = {};
        newValue[e.descriptor] = e.value;
        this.setState(newValue, () => { if (this.afterHandleValueChange) this.afterHandleValueChange(e); });
    },
};

module.exports = FormInputMixin;
