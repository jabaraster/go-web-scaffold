var React = require('react');
var Bootstrap = require('react-bootstrap');
var ButtonToolbar = Bootstrap.ButtonToolbar;
var ButtonGroup = Bootstrap.ButtonGroup;
var Button = Bootstrap.Button;
var Modal = Bootstrap.Modal;
var ModalTrigger = Bootstrap.ModalTrigger;
var Glyphicon = Bootstrap.Glyphicon

// Modalはコンポーネントで包まないとクローズボタンが動作しなくなる. 不可解・・・
var ModalSample = React.createClass({
    handle: function() {
        if (this.props.onClose() === false) {
            return;
        }
        this.props.onRequestHide(); // これを呼ぶとモーダルが消える. 不可解・・・
    },
    render: function() {
        return (
           <Modal {...this.props}
                  title={`モーダルダイアログ`}
                  animation={true}
                  backdrop={false}
           >
               <div className="modal-body">
                   本文
               </div>
               <div className="modal-footer">
                   <Button bsStyle="primary" onClick={this.handle}>Close</Button>
               </div>
           </Modal>
        );
    }
});

var Page = React.createClass({
    handleModalClose: function() {
        return confirm('閉じる？');
    },
    render: function() {
        // onRequestHideハンドラを指定しても動かない. 不可解・・・
        var modal = (
            <ModalSample onClose={this.handleModalClose} />
        );
        return (
            <div className="Page container">
                <h1>Bootstrap Components !</h1>
                <ButtonToolbar>
                    <Button bsStyle="default">Default</Button>
                    <Button bsStyle="primary">Primary</Button>
                    <Button bsStyle="success">Success</Button>
                    <Button bsStyle="warning">Warning</Button>
                    <Button bsStyle="danger">Danger</Button>
                </ButtonToolbar>
                <hr />
                <ButtonGroup>
                    {[1,2,3,4].map(function(e) { return (<Button bsStyle="default">{`${e}番目`}</Button>); })}
                </ButtonGroup>
                <ButtonGroup>
                    {[1,2,3,4,5,6].map(function(e) { return (<Button bsStyle="default">{`${e}番目`}</Button>); })}
                </ButtonGroup>
                <hr />
                <ModalTrigger modal={modal}>
                    <Button bsStyle="primary"><Glyphicon glyph="search" /> {`Show Dialog`}</Button>
                </ModalTrigger>
                <hr />
                <a href="" className="btn btn-default">
                    <Glyphicon glyph="refresh" />
                    リロード
                </a>
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
