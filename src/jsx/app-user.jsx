'use strict';

const request = require('./component/app-ajax.js');
const React = require('react');
const Bootstrap = require('react-bootstrap');
const Button = Bootstrap.Button;
const Glyphicon = Bootstrap.Glyphicon;
const Accordion = Bootstrap.Accordion;
const Panel = Bootstrap.Panel;
const Input = Bootstrap.Input;
const Modal = Bootstrap.Modal;
const Popover = Bootstrap.Popover;
const OverlayTrigger = Bootstrap.OverlayTrigger;
const InputField = require('./component/input-field.jsx');
const Message = require('./component/message.jsx');
const FormInputMixin = require('./mixin/form-input.js');

const AppUserEditor = React.createClass({
    mixins: [FormInputMixin],

    propTypes: {
        initialValues: React.PropTypes.object,
        onSave: React.PropTypes.func.isRequired
    },
    getInitialState: function() {
        return {
            userId: this.props.initialValues ? this.props.initialValues.UserId : '',
            password: '',
            passwordConfirmation: '',
            messages: []
        };
    },
    resetValues: function() {
        // TODO 今の作りではInputFieldに値を与えることができない.
        if (!this.props.initialValues) {
            return;
        }
        this.setState({
            userId: this.props.initialValues ? this.props.initialValues.UserId : '',
            password: '',
            passwordConfirmation: ''
        });
    },
    save: function() {
        if (this.hasError()) {
            return;
        }
        if (this.state.password !== this.state.passwordConfirmation) {
            this.setState({ messages: ['パスワードが一致していません.'] });
            return;
        }
        const e = {
            userId: this.state.userId,
            password: this.state.password,
            passwordConfirmation: this.state.passwordConfirmation,
        };
        request.post('/app-user/').
            type('form').
            send(e).
        end((err, res) => {
            if (err) {
                this.setState({ messages: [err] });
                return;
            }
            if (res.body.error) {
                this.setState({ messages: [res.body.error] });
                return;
            }
            this.setState({ messages: [] });
            this.props.onSave(e);
        });
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
                <Message messages={this.state.messages} />
            </div>
        );
    }
});

const AppUserEditorDialog = React.createClass({
    propTypes: {
        show: React.PropTypes.bool.isRequired,
        onSave: React.PropTypes.func.isRequired,
        onCancel: React.PropTypes.func.isRequired
    },
    render: function() {
        return (
            <Modal show={this.props.show} onHide={this.props.onCancel}>
                <Modal.Header closeButton>
                    <Modal.Title>ユーザ新規登録</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <AppUserEditor onSave={this.props.onSave} />
                </Modal.Body>
            </Modal>
        );
    }
});

const AppUserList = React.createClass({
    propTypes: {
    },
    getInitialState: function() {
        return {
            appUsers: [],
            openDialog: false
        };
    },
    componentDidMount: function() {
        request.get('/app-user/').
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
        this.setState({ openDialog: false });
    },
    render: function() {
        const rows = this.state.appUsers.map((appUser, idx) => {
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

const Page = React.createClass({
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
