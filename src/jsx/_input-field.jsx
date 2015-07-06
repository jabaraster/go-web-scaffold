'use strict';

var $ = require('jquery');
var React = require('react');
var Bootstrap = require('react-bootstrap');
var Input = Bootstrap.Input;
var Popover = Bootstrap.Popover;
var OverlayTrigger = Bootstrap.OverlayTrigger;

var InputField = React.createClass({
    propTypes: {
        required: React.PropTypes.bool.isRequired,
        placeholder: React.PropTypes.string.isRequired,
        label: React.PropTypes.string.isRequired,
        maxCharCount: React.PropTypes.number.isRequired,
        type: React.PropTypes.oneOf(['text', 'password', 'textarea']).isRequired,
        initialValue: React.PropTypes.string.isRequired,
        descriptor: React.PropTypes.string.isRequired,
        focus: React.PropTypes.bool,
        onValueChange: React.PropTypes.func.isRequired
    },
    getInitialState: function() {
        return {
            error: '',
            value: ''
        };
    },
    componentDidMount: function() {
        var err = '';
        if (this.props.required === true && this.props.initialValue.length === 0) {
            err = '必須入力です.';
        }
        this.setState({ target: this.refs.textField, value: this.props.initialValue, error: err }, () => {
            this.props.onValueChange({
                targetRef: this.refs.textField,
                value: this.props.initialValue,
                descriptor: this.props.descriptor,
                error: this.state.error
            });
        });
    },
    handleValueChange: function(e) {
        var newValue = e.target.value;
        if (this.props.required === true) {
            if (newValue.length === 0) {
                this.setState({ error: '必須入力です.', value: newValue }, () => {
                    this.props.onValueChange({
                        targetRef: this.refs.textField,
                        value: newValue,
                        descriptor: this.props.descriptor,
                        error: this.state.error
                    });
                });
                return;
            }
        }
        if (newValue.length > this.props.maxCharCount) {
            this.setState({ error: '' });
            return;
        }
        this.setState({ value: newValue }, () => {
            this.setState({ error: '', value: newValue }, () => {
                this.props.onValueChange({ value: newValue, descriptor: this.props.descriptor, error: this.state.error });
            });
        });
    },
    render: function() {
        return (
            <OverlayTrigger
                trigger="focus"
                placement='right'
                overlay={<Popover title="">{this.state.error?this.state.error:'OK!'}</Popover>}
            >
                <Input type={this.props.type}
                       label={this.props.label}
                       placeholder={this.props.placeholder}
                       value={this.state.value}
                       ref="textField"
                       onChange={this.handleValueChange}
                       onFocus={e => this.setState({ target: e.target })}
                       className={this.state.error?'error':''}
                />
            </OverlayTrigger>
        );
    }
});

module.exports = InputField;
