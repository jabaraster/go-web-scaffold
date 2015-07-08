'use strict';

const React = require('react');
const Fluxxor = require('fluxxor');

const constants = {
    UPDATE_COUNTER: 'UPDATE_COUNTER',
    UPDATE_TIME: 'UPDATE_TIME',
};

// Store
const CounterStore = Fluxxor.createStore({
    initialize: function() {
        this.counter = 0;
        this.time = new Date();
        this.bindActions(
            constants.UPDATE_COUNTER, this.onUpdateCoutner
        );
        this.bindActions(
            constants.UPDATE_TIME, this.onUpdateTime
        );
    },
    onUpdateCoutner: function(payload) {
        this.counter = this.counter + payload.value;
        this.emit('change');
    },
    onUpdateTime: function(payload) {
        this.time = new Date();
        this.emit('change');
    },
    getState: function() {
        return {
            counter: this.counter,
            time: this.time,
        };
    }
});

// Action(Action Creator)
const actions = {
    plusCounter: function() {
        this.dispatch(constants.UPDATE_COUNTER, { value: 1 });
    },
    minusCounter: function() {
        this.dispatch(constants.UPDATE_COUNTER, { value: -1 });
    },
    updateTime: function() {
        this.dispatch(constants.UPDATE_TIME, {});
    },
};

// React‚©‚ç—˜—p‚·‚éMixin
const FluxMixin = Fluxxor.FluxMixin(React);
const StoreWatchMixin = Fluxxor.StoreWatchMixin;

// View(React)
const CounterApp = React.createClass({
    mixins: [FluxMixin, StoreWatchMixin('CounterStore')],

    getStateFromFlux: function() {
        return this.getFlux().
                    store('CounterStore').
                    getState();
    },
    render: function() {
        var time = this.state.time;
        return (
            <div className="Page container">
                <Counter value={this.state.counter} />
                <div>
                    <span>Time: {`${time.getHours()}:${time.getMinutes()}:${time.getSeconds()}.${time.getMilliseconds()}`}</span>
                </div>
            </div>
        );
    }
});

const Counter = React.createClass({
    mixins: [FluxMixin],

    propTypes: {
        value: React.PropTypes.number.isRequired,
    },
    onClickPlus: function() {
        return this.getFlux().actions.plusCounter();
    },
    onClickMinus : function() {
        return this.getFlux().actions.minusCounter();
    },
    onClickTime: function() {
        return this.getFlux().actions.updateTime();
    },
    render: function() {
        return (
            <div>
                <span>count: {this.props.value}</span>
                <div>
                    <button onClick={this.onClickPlus}>+1</button>
                    <button onClick={this.onClickMinus}>-1</button>
                </div>
                <button onClick={this.onClickTime}>Update Time</button>
            </div>
        );
    }
});

const stores = { CounterStore: new CounterStore() };
const flux = new Fluxxor.Flux(stores, actions);

React.render(
    <CounterApp flux={flux} />,
    document.getElementById('page')
);
