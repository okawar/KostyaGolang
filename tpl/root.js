function Main() {
    this.__proto__ = new Util();
    const util = this.__proto__;

    this.init = () => {
        console.log("init");
        this.hideAll(); // Сначала скрываем все элементы
        document.querySelectorAll(".action").forEach(el => {
            console.log(el);
            el.onclick = () => {
                this[el.dataset.action]();
            };
        });
        this.checkTokenValidity();
    };

    this.hideAll = () => {
        util.id("authForm").style.display = "none";
        util.id("studentsTable").style.display = "none";
        util.id("addStudentBtn").style.display = "none";
        util.id("deleteStudentBtn").style.display = "none";
        util.id("logoutBtn").style.display = "none";
        util.id("CourseTable").style.display = "none";
        util.id("addCourseBtn").style.display = "none";
        util.id("deleteCourseBtn").style.display = "none";
        document.getElementById('paymentInfo').style.display = "none";
        document.getElementById('showPaymentInfoBtn').style.display = "none";
    };


    this.checkTokenValidity = () => {
        const token = localStorage.getItem("token");
        util.get("/validateToken?token=" + token, data => {
            console.log(data, "Token validation response");
            if (data == true) {
                this.showInterface();
            } else {
                this.showAuthForm();
            }
        });
    };

    this.showAuthForm = () => {
        util.id("authForm").style.display = "block";
    };

 /*   this.showStudentsTable = () => {
        util.id("authForm").style.display = "none";
        console.log("дошло не дошло")
        util.id("studentsTable").style.display = "block";
        util.id("addStudentBtn").style.display = "inline-block";
        util.id("deleteStudentBtn").style.display = "inline-block";
        util.id("logoutBtn").style.display = "block";
        this.getStudents(); // Получаем и отображаем пользователей
    };
*/

    this.showInterface = () => {
        util.id("authForm").style.display = "none";
        util.id("studentsTable").style.display = "block";
        util.id("addStudentBtn").style.display = "inline-block";
        util.id("deleteStudentBtn").style.display = "inline-block";
        util.id("logoutBtn").style.display = "block";
        util.id("CourseTable").style.display = "block";
        util.id("addCourseBtn").style.display = "inline-block";
        util.id("deleteCourseBtn").style.display = "inline-block";
        this.getStudents(); // Получаем и отображаем пользователей
        this.getCourse()
        document.getElementById('showPaymentInfoBtn').style.display = 'block';
        document.getElementById('showPaymentInfoBtn').addEventListener('click', function() {
            var paymentInfo = document.getElementById('paymentInfo');
            paymentInfo.style.display = (paymentInfo.style.display === 'none' ? 'block' : 'none');
        });
    };

    this.hideInterface = () => {
        util.id("authForm").style.display = "none";
        util.id("studentsTable").style.display = "none";
        util.id("addStudentBtn").style.display = "none";
        util.id("deleteCourseBtn").style.display = "none";
        util.id("deleteCourseBtn").style.display = "none";
        util.id("logoutBtn").style.display = "block";
        util.id("CourseTable").style.display = "block";
        util.id("addCourseBtn").style.display = "none";
        this.getCourse()
        document.getElementById('showPaymentInfoBtn').style.display = 'block';
        document.getElementById('showPaymentInfoBtn').addEventListener('click', function() {
            var paymentInfo = document.getElementById('paymentInfo');
            paymentInfo.style.display = (paymentInfo.style.display === 'none' ? 'block' : 'none');
        });
    };

    this.addStudent = () => {
        util.modal('addStudentModal', 'show');

        document.getElementById('saveStudentBtn').onclick = () => {
            const student = {
                login: util.id('addStudentLogin').value,
                pass: util.id('addStudentPassword').value,
                name: util.id('addStudentName').value,
                access: util.id('addStudentAccess').value === 'true'
            };

            util.post('/addStudent', student, (response) => {
                if (response.success) {
                    this.getStudents();
                } else {
                    alert(response.message);
                }
            });

            util.modal('addStudentModal', 'hide');
        };
    };


    this.deleteStudent = () => {
        // Вызываем модальное окно
        util.modal('confirmDeleteModal', 'show');

        // Получаем кнопку подтверждения удаления и добавляем обработчик событий
        document.getElementById('confirmDeleteBtn').onclick = () => {
            const studentId = parseInt(document.getElementById('deleteStudentIdInput').value, 10);

            util.post('/deleteStudent', {sid: studentId}, (response) => {
                // Обработка ответа сервера
                if (response.success) {
                    // Обновить таблицу студентов
                    this.getStudents();
                } else {
                    // Ошибка при удалении студента
                    alert(response.message);
                }
            });

            // Закрыть модальное окно
            util.modal('confirmDeleteModal', 'hide');
        };
    };


    this.getStudents = () => {
        util.get("/getStudentsTable", data1 => {
            console.log(data1, "Data from /getStudentsTable");

            if (Array.isArray(data1)) {
                const tableBody = util.id("studentsTable").querySelector("tbody");
                tableBody.innerHTML = ""; // Очищаем текущие строки

                data1.forEach(student => {
                    const row = `<tr>
                    <td>${student.sid}</td>
                    <td>${student.Login}</td>
                    <td>${student.Pass}</td>
                    <td>${student.Name}</td>
                    <td>${student.Access ? 'Да' : 'Нет'}</td>  
                </tr>`;
                    tableBody.innerHTML += row;
                });
            } else {
                console.error("Data is not an array:", data1);
            }
        }, error => {
            console.error("Error fetching data:", error);
        });
    };

    this.addCourse = () => {
        util.modal('addCourseModal', 'show');

        document.getElementById('saveCourseBtn').onclick = () => {
            const vds = {
                name: util.id('addCourseName').value,
                description: util.id('addCourseDescription').value,
                price: parseInt(util.id('addCoursePrice').value),
                course_holder: util.id('addCourse_Holder').value,
            };

            util.post('/addCourse', vds, (response) => {
                if (response.success) {
                    console.log("Course added succefuly")
                    this.getCourse();
                } else {
                    console.error('Error adding Course', response.message)
                    alert(response.message);
                }
            });

            util.modal('addCourseModal', 'hide');
        };
    };

    this.getCourse = () => {
        util.get("/getCourseTable", data2 => {
            console.log(data2, "Data from /getCourseTable");

            if (Array.isArray(data2)) {
                const tableBody = util.id("CourseTable").querySelector("tbody");
                tableBody.innerHTML = ""; // Очищаем текущие строки

                data2.forEach(vds => {
                    const row = `<tr>
                <td>${vds.uid}</td>
                <td>${vds.Name}</td>
                <td>${vds.Description}</td>
                <td>${vds.Price}</td>
                <td>${vds.Course_Holder}</td>
            </tr>`;
                    tableBody.innerHTML += row;
                });
            } else {
                console.error("Data is not an array:", data2);
            }
        }, error => {
            console.error("Error fetching data:", error);
        });
    };
    this.deleteCourse = () => {
        // Вызываем модальное окно
        util.modal('confirmDeleteModalCourse', 'show');

        // Получаем кнопку подтверждения удаления и добавляем обработчик событий
        document.getElementById('confirmDeleteCourseBtn').onclick = () => {
            const vdsId = parseInt(document.getElementById('deleteCourseIdInput').value, 10);

            util.post('/deleteCourse', {vid: vdsId}, (response) => {

                // Обработка ответа сервера
                if (response.success) {
                    // Обновить таблицу пользователей
                    this.getCourse();
                } else {
                    // Ошибка при удалении пользователя
                    alert(response.message);
                }
            });

            // Закрыть модальное окно
            util.modal('confirmDeleteModalCourse', 'hide');
        };
    };

    this.authIn = () => {
        //console.log("flskndlkfdsn11")
        util.post("/auth", {
            login: util.id("authLogin").value,
            pass: util.id("authPassword").value
        }, data => {
            //console.log(data,"qqq")
            if (data["token"]) {
                //console.log(data,"qqqw")
                localStorage.setItem("token", data["token"])
                util.setToken(data["token"]);

                if (data["access"]) {
                    this.showInterface();
                    document.getElementById('showPaymentInfoBtn').style.display = 'block';

                } else {
                    this.hideInterface()
                }
            } else {
                alert("Ytf")
            }
        })
    };

    this.init();
}

    document.getElementById('logoutBtn').addEventListener('click', function() {
    fetch('/logout', {
        method: 'POST',
        headers: {
            'Authorization': localStorage.getItem('token')
        }
    })
        .then(response => response.text())
        .then(data => {
            console.log(data);
            localStorage.removeItem('token');
            window.location.reload();
        })
        .catch(error => console.error('Ошибка:', error));
});
document.getElementById('deleteStudentBtn').addEventListener('click', function() {
    main.deleteStudent();
});
document.getElementById('deleteCourseBtn').addEventListener('click', function() {
    main.deleteCourse();
});

const main = new Main();
