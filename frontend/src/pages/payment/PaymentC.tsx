import React, { useState, useEffect } from 'react';
import {
    Space, Button, Col, Row, Divider, Form, Input, Card,
    message, DatePicker, InputNumber, Upload
} from "antd";
import { PlusOutlined, UploadOutlined } from "@ant-design/icons";
import { PaymentInterface } from "../../interfaces/IPayment";
import { CreatePayment, GetReservesByShopId } from "../../services/https/index";
import { useNavigate, Link, useLocation } from "react-router-dom";
import { ReserveDetailsInterface } from '../../interfaces/IReserveDetails';

function PaymentC() {
    const navigate = useNavigate();
    const location = useLocation();
    const [messageApi, contextHolder] = message.useMessage();
    const [slipFileName, setSlipFileName] = useState<string>("");
    const [form] = Form.useForm();
    

    // Extract reserveId from URL query params
    const query = new URLSearchParams(window.location.search);
  const reserveId = query.get('reserveId');
  const [reserve, setReserve] = useState<ReserveDetailsInterface>();

  const onFinish = async (values: PaymentInterface) => {
    // แปลง reserveId จาก string เป็น number และตรวจสอบว่ามีค่า
    const reserveIdNumber = reserveId ? parseInt(reserveId) : null;

    if (reserveIdNumber !== null) {
      try {
        // เรียก API เพื่อดึงข้อมูลการจอง
        let res = await GetReservesByShopId(reserveIdNumber);
        
        console.log("Full API Response:", res);  // ดูข้อมูลเต็มที่ API ส่งกลับมา
        console.log("Response data:", res.data); // ดูเฉพาะข้อมูลที่อยู่ใน data
        
        setReserve(res.data);

        // ตรวจสอบว่าข้อมูลที่ส่งกลับมาเป็น array หรือไม่ และตรวจสอบว่ามีข้อมูลใน array
        if (Array.isArray(res.data) && res.data.length > 0) {
          const reserveIdValue = res.data[0].id; // เข้าถึงค่า id จาก array

          // ตรวจสอบและจัดการข้อมูลให้ตรงตามโครงสร้างที่ต้องการ
          const valuesWithReserveId = {
            Name: values.name,
            Date: values.date,
            TotalPrice: values.total_price,
            Slip: values.slip,
            ReserveID: reserveIdValue // ใช้ค่า reserveId จากข้อมูลที่ส่งกลับมา
          };

          try {
            // เรียก API เพื่อสร้างการชำระเงิน โดยส่ง valuesWithReserveId ที่มี reserveId
            const resPayment = await CreatePayment(valuesWithReserveId);

            if (resPayment.status === 201) {
              message.success(resPayment.data.message);
              setTimeout(() => {
                navigate(`/PaymentC?reserveId=${reserveId}`);
              }, 2000);
            } else {
              message.error(resPayment.data.error);
            }
          } catch (error) {
            message.error("เกิดข้อผิดพลาดในการสร้างการชำระเงิน");
          }
        } else {
          message.error("ไม่พบข้อมูลการจองหรือข้อมูลไม่สมบูรณ์");
        }
      } catch (error) {
        message.error("เกิดข้อผิดพลาดในการดึงข้อมูลการจอง");
      }
    } else {
      message.error("reserveId ไม่ถูกต้อง");
    }
  };


    

    // Handle image upload
    const handleUpload = (file: any) => {
        const reader = new FileReader();

        reader.onload = () => {
            const base64String = reader.result as string; // Convert to base64 string
            form.setFieldsValue({ slip: base64String }); // Save base64 string to form field
            setSlipFileName(file.name); // Set file name
        };

        reader.readAsDataURL(file); // Read file as base64 string
        return false; // Prevent default upload behavior
    };

    return (
        <div style={styles.container}>
            {contextHolder}
            <Card style={styles.card}>
                <h2>โปรดกรอกข้อมูลการชำระเงิน</h2>
                <Divider />

                <Form form={form} name="basic" layout="vertical" onFinish={onFinish} autoComplete="off">
                    <Row gutter={[16, 0]}>
                        <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                            <Form.Item
                                label="ชื่อบัญชี"
                                name="name"
                                rules={[{ required: true, message: "กรุณากรอกชื่อบัญชี !" }]}
                            >
                                <Input />
                            </Form.Item>
                        </Col>
                    </Row>

                    <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                        <Form.Item
                            label="วัน/เดือน/ปี ที่ชำระเงิน"
                            name="date"
                            rules={[{ required: true, message: "กรุณาเลือกวัน/เดือน/ปี ที่ชำระเงิน !" }]}
                        >
                            <DatePicker style={{ width: "100%" }} />
                        </Form.Item>
                    </Col>

                    <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                        <Form.Item
                            label="จำนวนเงิน"
                            name="total_price"
                            rules={[{ required: true, message: "กรุณากรอกจำนวนเงิน !" }]}
                        >
                            <InputNumber min={0} defaultValue={0} style={{ width: "100%" }} />
                        </Form.Item>
                    </Col>

                    {/* Upload Slip */}
                    <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                        <Form.Item
                            label="อัพโหลดหลักฐานการชำระเงิน"
                            name="slip"
                            rules={[{ required: true, message: "กรุณาอัพโหลดหลักฐานการชำระเงิน !" }]}
                        >
                            <Upload beforeUpload={handleUpload} listType="text" maxCount={1}>
                                <Button icon={<UploadOutlined />}>อัพโหลดหลักฐานการชำระเงิน</Button>
                            </Upload>
                            {slipFileName && <div>ชื่อไฟล์: {slipFileName}</div>}
                        </Form.Item>
                    </Col>

                    <Row justify="end">
                        <Col style={{ marginTop: "40px" }}>
                            <Form.Item>
                                <Space>
                                    <Link to="/Payments">
                                        <Button htmlType="button" style={{ marginRight: "10px" }}>
                                            ยกเลิก
                                        </Button>
                                    </Link>
                                    <Button type="primary" htmlType="submit" icon={<PlusOutlined />}>
                                        ยืนยัน
                                    </Button>
                                </Space>
                            </Form.Item>
                        </Col>
                    </Row>
                </Form>
            </Card>
        </div>
    );
}

export default PaymentC;

const styles = {
    container: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundColor: '#f0f2f5'
    },
    card: {
        width: '600px',
        padding: '20px',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    }
};
