import { useState } from 'react';
import './App.css'
import { Button, Divider, Form, Input, Space, Typography } from 'antd'
const { Title } = Typography;

function App() {
  const [array, setArray] = useState('');
  const [target, setTarget] = useState('');
  const [output, setOutput] = useState('');

  const handleSearch = () => {
    const arrayData = array.split(',').map(Number);
    const targetData = Number(target);

    if (arrayData.some(isNaN)) {
      setOutput('Массив должен содержать только числа');
      return;
    }

    fetch('http://localhost:8080/search', {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        array: arrayData,
        target: targetData,
      }),
    })
      .then(response => response.json())
      .then(data => {
        setOutput(`${data.result}`);
      })
      .catch(error => {
        console.error(error);
        setOutput('Произошла ошибка при выполнении поиска');
      });
  };

  return (
    <div className='container'>
      <Space direction='vertical'>
        <Typography>
          <Title>
            Троичный(тернарный) поиск
          </Title>
          <Divider />
        </Typography>
        <Form>
          <Form.Item name="ArrayInput" rules={[{ required: true, message: 'Введите отсортированный массив данных' }]}>
            <Input   size="large" placeholder='Введите массив данных через запятую' onChange={e => setArray(e.target.value)} />
          </Form.Item>
          <Form.Item name="TargetInput" rules={[{ required: true, message: 'Введите число которое нужно найти' }]}>
            <Input  size="large" placeholder='число которое надо найти' onChange={e => setTarget(e.target.value)} />
          </Form.Item>
          <Form.Item>
            <Button  htmlType="submit" onClick={handleSearch}>Найти</Button>
          </Form.Item>
        </Form>
        <Space>
        <div>
          {(() => {
            if (output == "-1") {
              return <Title level={4}>
              Вывод: нет в массиве
              </Title>;
            } else {
              return <Title level={4}>
              Вывод: элемент {target} имет индекс {output}
              </Title>;
            }
          })()}
        </div>
        </Space>
      </Space>
    </div>
  );
}

export default App;
