'use strict';

var request = require('superagent');
var React = require('react');
var Bootstrap = require('react-bootstrap');
var Button = Bootstrap.Button;
var Input = Bootstrap.Input;
var Popover = Bootstrap.Popover;
var OverlayTrigger = Bootstrap.OverlayTrigger;
var InputField = require('./_input-field.jsx');

var Page = React.createClass({
    getInitialState: function() {
        return {
            errors: {},
            userId: '',
            password: ''
        };
    },
    hasError: function() {
        for (var p in this.state.errors) {
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
        this.state[e.descriptor] = e.value;
        this.setState(this.state);
    },
    handleLoginClick: function() {
        if (this.hasError()) {
            return;
        }
        request.post('/model/authenticator').
            type('form').
            send({ userId: this.state.userId, password: this.state.password }).
            end((err, res) => {
                if (!res.body.error) {
                    location.href = '/';
                }
            });
    },
    render: function() {
        return (
            <div className="Page container">
                <InputField label="ユーザID"
                            required={true}
                            placeholder="メールアドレス形式"
                            maxCharCount={100}
                            type="text"
                            initialValue={this.state.userId}
                            descriptor="userId"
                            focus={true}
                            onValueChange={this.handleValueChange}
                />
                <InputField label="ユーザID"
                            required={true}
                            placeholder=""
                            maxCharCount={100}
                            type="password"
                            initialValue={this.state.password}
                            descriptor="password"
                            onValueChange={this.handleValueChange}
                />
                <Button bsStyle="success" onClick={this.handleLoginClick}>ログイン</Button>
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
