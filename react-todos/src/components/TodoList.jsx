import React, { useEffect, useState, useCallback } from "react";
import { Tabs, Layout, Row, Col, Input, message } from "antd";
import "./TodoList.css";
import TodoTab from "./TodoTab";
import TodoForm from "./TodoForm";
import { createTodo, deleteTodo, getAllTodos, updateTodo } from "../services/todosService";
const { TabPane } = Tabs;
const { Content } = Layout;

const TodoList = () => {
    const [refreshing, setRefreshing] = useState(false);
    const [todos, setTodos] = useState();
    const [activeTodos, setActiveTodos] = useState();
    const [completedTodos, setCompletedTodos] = useState();

    const handleFormSubmit = useCallback(async (todo) => {
        console.log('handleFormSubmit', todo);
        createTodo(todo).then(onRefresh());
        message.success('Todo created successfully');
    }, []);

    const handleRemoveTodo = useCallback(async (todo) => {
        deleteTodo(todo.ID).then(onRefresh());
        message.warn('Todo deleted successfully');
    }, []);

    const handleToogleTodoStatus = useCallback(async (todo) => {
        todo.completed = !todo.completed;
        updateTodo(todo).then(onRefresh());
        message.info('Todo updated successfully');
    }, []);

    const refresh = () => {
        getAllTodos().then(todos => {
            setTodos(todos);
            setActiveTodos(todos.filter(todo => !todo.completed));
            setCompletedTodos(todos.filter(todo => todo.completed));
        }).then(console.log('refreshed'));
    }

    const onRefresh = useCallback(async () => {
        setRefreshing(true);
        let todos = await getAllTodos();
        setTodos(todos);
        setActiveTodos(todos.filter(todo => !todo.completed));
        setCompletedTodos(todos.filter(todo => todo.completed));
        setRefreshing(false);
        console.log('Refreshed state', refreshing);
    }, [refreshing]);

    useEffect(() => {
        refresh();
    }, [onRefresh]);

    return (
        <Layout className="layout">
            <Content className="content" style={{ padding: '0 50px' }}>
                <div className="todoList">
                    <Row>
                        <Col span={14} offset={5}>
                            <h1>HansMelo Todos</h1>
                            <TodoForm onFormSubmit={handleFormSubmit} />
                            <br />
                            <Tabs defaultActiveKey="all">
                                <TabPane tab="All" key="all">
                                    <TodoTab todos={todos} onTodoToggle={handleToogleTodoStatus} onTodoRemoval={handleRemoveTodo} />
                                </TabPane>
                                <TabPane tab="Active" key="active">
                                    <TodoTab todos={activeTodos} onTodoToggle={handleToogleTodoStatus} onTodoRemoval={handleRemoveTodo} />
                                </TabPane>
                                <TabPane tab="Completed" key="completed">
                                    <TodoTab todos={completedTodos} onTodoToggle={handleToogleTodoStatus} onTodoRemoval={handleRemoveTodo} />
                                </TabPane>
                            </Tabs>
                        </Col>
                    </Row>
                </div>
            </Content>
        </Layout>
    );
}

export default TodoList;