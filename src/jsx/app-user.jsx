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
    componentDidMount: function() {
        var err = '';
        if (this.props.required === true && this.props.initialValue.length === 0) {
            err = '必須入力です.';
        }
        this.setState({ target: this.refs.textField, error: err }, () => {
            this.props.onValueChange({ value: this.props.initialValue, descriptor: this.props.descriptor, error: this.state.error });
        });
    },
    handleValueChange: function(e) {
        var newValue = e.target.value;
        if (this.props.required === true) {
            if (newValue.length === 0) {
                this.setState({ error: '必須入力です.', value: newValue }, () => {
                    this.props.onValueChange({ value: newValue, descriptor: this.props.descriptor, error: this.state.error });
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
            passwordConfirmation: '',
            inputContexts: {}
        };
    },
    handleValueChange: function(e) {
        var ctx = this.state.inputContexts[e.descriptor];
        if (!ctx) {
            ctx = {};
            this.state.inputContexts[e.descriptor] = ctx;
        }
        const hasError = !!e.error;
        ctx.error = e.error;
        if (!e.error) {
            ctx.value = e.value;
        }
        this.setState({ inputContexts: this.state.inputContexts }, () => console.log(this.state.inputContexts));
    },
    hasError: function() {
        const errors = this.state;
        for(var e in errors) {
            if (errors[e] === true) {
                return true;
            }
        }
        return false;
    },
    save: function() {
        if (this.hasError()) {
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
                                placeholder="半角英数記号"
                                maxCharCount={100}
                                type="password"
                                initialValue={this.state.password}
                                descriptor="password"
                                onValueChange={this.handleValueChange}
                    />
                    <InputField label="パスワード確認"
                                required={true}
                                placeholder="半角英数記号"
                                maxCharCount={100}
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
