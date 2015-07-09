'use strict';

const request = require('./component/app-ajax.js');
const React = require('react');
const Bootstrap = require('react-bootstrap');
const Button = Bootstrap.Button;
const Input = Bootstrap.Input;
const Popover = Bootstrap.Popover;
const OverlayTrigger = Bootstrap.OverlayTrigger;
const InputField = require('./component/input-field.jsx');
const Glyphicon = Bootstrap.Glyphicon
const Message = require('./component/message.jsx');
const FormInputMixin = require('./mixin/form-input.js');

const Page = React.createClass({
    mixins: [FormInputMixin],

    getInitialState: function() {
        return {
            userId: '',
            password: '',
            messages: []
        };
    },
    handleLoginClick: function() {
        if (this.hasError()) {
            this.setState({ messages: ['入力内容に誤りがあります.'] });
            return;
        }
        request.post('/authenticator').
            type('form').
            send({ userId: this.state.userId, password: this.state.password }).
            end((err, res) => {
                if (err) {
                    this.setState({ messages: [err] });
                    return;
                }
                if (res.body.error) {
                    this.setState({ messages: [res.body.error] });
                    return;
                }
                // ログイン成功時の処理
                location.href = '/';
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
                <Button bsStyle="success" onClick={this.handleLoginClick}>
                    <Glyphicon glyph="log-in" />
                    ログイン
                </Button>
                <Message messages={this.state.messages} />
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
