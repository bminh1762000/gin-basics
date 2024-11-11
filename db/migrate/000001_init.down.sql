ALTER TABLE "Task_Categories" DROP CONSTRAINT "Task_Categories_category_id_fkey";
ALTER TABLE "Task_Categories" DROP CONSTRAINT "Task_Categories_task_id_fkey";
ALTER TABLE "Categories" DROP CONSTRAINT "Categories_user_id_fkey";
ALTER TABLE "Tasks" DROP CONSTRAINT "Tasks_user_id_fkey";

DROP TABLE IF EXISTS "Categories";

DROP TABLE IF EXISTS "Users";

DROP TABLE IF EXISTS "Tasks";

DROP TABLE IF EXISTS "Task_Categories";
