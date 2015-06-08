(function($) {
"use strict";

window.DateTimeInput = React.createClass({displayName: "DateTimeInput",
    propTypes: {
        placeholder: React.PropTypes.string.isRequired
    },
    componentDidMount: function() {
        $(this.getDOMNode()).datetimepicker({
            lang: 'ja',
            inline: false,
            step: 15,
            //theme: 'dark',
            showButtonPanel: true
        });
    },
    render: function() {
        return (
            React.createElement("input", {type: "text", className: "DateTimeInput form-control", placeholder: this.props.placeholder})
        );
    }
});

window.DateInput = React.createClass({displayName: "DateInput",
    propTypes: {
        placeholder: React.PropTypes.string.isRequired
    },
    componentDidMount: function() {
        $(this.getDOMNode()).datepicker({
            showButtonPanel: true
        });
    },
    render: function() {
        return (
            React.createElement("input", {type: "text", className: "DateInput form-control", placeholder: this.props.placeholder})
        );
    }
});

window.Input = React.createClass({displayName: "Input",
    propTypes: {
        placeholder: React.PropTypes.string.isRequired,
        maxLength: React.PropTypes.number.isRequired,
        required: React.PropTypes.bool.isRequired,
        value: React.PropTypes.string
    },
    getInitialState: function() {
        return {
            value: this.props.value ? this.props.value : ''
        };
    },
    handleChange: function(e) {
        if (e.target.value.length > this.props.maxLength) {
            return;
        }
        this.setState({ value: e.target.value });
    },
    render: function() {
        var cls = React.addons.classSet({
            'form-control': true,
            'error': this.props.required && this.state.value.length === 0
        });
        return (
            React.createElement("div", {className: "Input"}, 
                React.createElement("input", {type: "text", 
                       className: cls, 
                       placeholder: this.props.placeholder, 
                       value: this.state.value, 
                       onChange: this.handleChange}
                ), 
                React.createElement("span", null, "残り", this.props.maxLength - this.state.value.length, "文字")
            )
        );
    }
});

})(jQuery);
