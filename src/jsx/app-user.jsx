'use strict';

var request = require('superagent');
var React = require('react');
var Bootstrap = require('react-bootstrap');
var Button = Bootstrap.Button;
var Glyphicon = Bootstrap.Glyphicon;
var Accordion = Bootstrap.Accordion;
var Panel = Bootstrap.Panel;
var Input = Bootstrap.Input;
var Modal = Bootstrap.Modal;
var Popover = Bootstrap.Popover;
var OverlayTrigger = Bootstrap.OverlayTrigger;
var InputField = require('./_input-field.jsx');

var AppUserEditor = React.createClass({
    propTypes: {
        initialValue: React.PropTypes.object,
        onSave: React.PropTypes.func.isRequired
    },
    getInitialState: function() {
        return {
            userId: '',
            userIdError: '',
            password: '',
            passwordConfirmation: ''
        };
    },
    componentDidMount: function() {
        this.resetValues();
    },
    handleValueChange: function(e) {
        var s = {};
        s[e.descriptor] = e.value;
        this.setState(s);
    },
    handleTextValueChange: function(propertyName, e) {
        console.log(arguments);
        var s = {};
        s[propertyName] = e.target.value;
        this.setState(s);
    },
    resetValues: function() {
        if (!this.props.initialValue) {
            return;
        }
        this.setState({
            userId: this.props.initialValues.UserId,
            password: '',
            passwordConfirmation: ''
        });
    },
    save: function() {
        var hasError = false;
        if (!this.state.appUserUserId) {
            this.setState({ hasUserIdError: true, userIdError: '必須入力です.' });
            hasError = true;
        }
        if (!this.state.password) {
        }

        if (hasError) {
            return;
        }
        // TODO
        var e = {
            UserId: this.state.userId,
            Password: this.state.password,
            PasswordConfirmation: this.state.passwordConfirmation
        };
        this.props.onSave(e);
    },
    render: function() {
        return (
            <div className="AppUserEditor">
                <div className="form-group">
                    <InputField label="ユーザID"
                                required={true}
                                placeholder="メールアドレス形式"
                                maxCharCount={100}
                                type="text"
                                initialValue={this.state.userId}
                                descriptor="userId"
                                onValueChange={this.handleValueChange}
                    />
                    <InputField label="パスワード"
                                required={true}
                                placeholder="英数記号"
                                maxCharCount={32}
                                type="password"
                                initialValue={this.state.password}
                                descriptor="password"
                                onValueChange={this.handleValueChange}
                    />
                    <InputField label="パスワード確認"
                                required={true}
                                placeholder="英数記号"
                                maxCharCount={32}
                                type="password"
                                initialValue={this.state.passwordConfirmation}
                                descriptor="passwordConfirmation"
                                onValueChange={this.handleValueChange}
                    />
                </div>
                <div className="form-group">
                    <Button bsStyle="primary" onClick={this.resetValues}><Glyphicon glyph="refresh" /> 初期値に戻す</Button>
                    <Button bsStyle="success" onClick={this.save}><Glyphicon glyph="ok" /> 保存</Button>
                </div>
            </div>
        );
    }
});

var AppUserEditorDialog = React.createClass({
    propTypes: {
        show: React.PropTypes.bool.isRequired,
        onSave: React.PropTypes.func.isRequired,
        onCancel: React.PropTypes.func.isRequired
    },
    render: function() {
        return (
            <Modal show={this.props.show} onHide={this.props.onCancel}>
                <Modal.Header closeButton>
                    <Modal.Title>Modal heading</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <AppUserEditor onSave={this.props.onSave} />
                </Modal.Body>
            </Modal>
        );
    }
});

var AppUserList = React.createClass({
    propTypes: {
    },
    getInitialState: function() {
        return {
            appUsers: [],
            openDialog: false
        };
    },
    componentDidMount: function() {
        request.get('/resource/app-user/').
            end((err, res) => {
                this.setState({ appUsers: res.body });
            });
    },
    showInputDialog: function() {
        this.setState({ openDialog: true });
    },
    closeInputDialog: function() {
        this.setState({ openDialog: false });
    },
    saveNewUser: function(e) {
        console.log(e);
        this.setState({ openDialog: false });
    },
    render: function() {
        var rows = this.state.appUsers.map((appUser, idx) => {
            return (
                <Panel header={appUser.UserId} eventKey={idx} key={appUser.ID}>
                    <AppUserEditor initialValues={appUser} onSave={this.saveNewUser} />
                </Panel>
            )
        });
        return (
            <div className="AppUserList">
                <Button onClick={this.showInputDialog} bsStyle="primary"><Glyphicon glyph="plus" /></Button>
                <Accordion className="app-user-row">
                    {rows}
                </Accordion>
                <AppUserEditorDialog show={this.state.openDialog} onCancel={this.closeInputDialog} onSave={this.saveNewUser} />
            </div>
        );
    }
});

var Page = React.createClass({
    render: function() {
        return (
            <div className="Page container">
                <h1>ユーザ一覧</h1>
                <AppUserList />
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
