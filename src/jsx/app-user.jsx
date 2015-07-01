const request = require('superagent');
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

const InputField = React.createClass({
    propTypes: {
        required: React.PropTypes.bool.isRequired,
        placeholder: React.PropTypes.string.isRequired,
        label: React.PropTypes.string.isRequired,
        maxCharCount: React.PropTypes.number.isRequired,
        type: React.PropTypes.oneOf(['text', 'password', 'textarea']).isRequired,
        initialValue: React.PropTypes.string.isRequired,
        descriptor: React.PropTypes.string.isRequired,
        onValueChange: React.PropTypes.func.isRequired
    },
    getInitialState: function() {
        return {
            error: '',
            value: ''
        };
    },
    handleValueChange: function(e) {
        var newValue = e.target.value;
        if (this.props.required === true) {
            if (newValue.length === 0) {
                this.setState({ error: '必須入力です.', value: newValue });
                return;
            }
        }
        if (newValue.length > this.props.maxCharCount) {
            this.setState({ error: '' });
            return;
        }
        this.setState({ value: newValue }, () => {
            this.setState({ error: '', value: newValue }, () => {
                this.props.onValueChange({ value: newValue, descriptor: this.props.descriptor });
            });
        });
    },
    render: function() {
        return (
            <OverlayTrigger trigger='focus' overlay={<Popover show={!!this.state.error} placement="left" title="">{this.state.error}</Popover>}>
                <Input type={this.props.type}
                       label={this.props.label}
                       placeholder={this.props.placeholder}
                       value={this.state.value}
                       onChange={this.handleValueChange}
                />
            </OverlayTrigger>
        );
    }
});

const AppUserEditor = React.createClass({
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
        const s = {};
        s[e.descriptor] = e.value;
        this.setState(s);
    },
    handleTextValueChange: function(propertyName, e) {
        console.log(arguments);
        const s = {};
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
                    <Modal.Title>Modal heading</Modal.Title>
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
        request.get('/model/app-user/').
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
